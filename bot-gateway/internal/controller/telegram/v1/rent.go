package v1

import (
	"context"
	"fmt"
	"gateway/internal/controller/telegram/dto"
	"gateway/internal/domain/books/model"
	rentService "gateway/internal/domain/rent/dto"
	"gateway/pkg/adapters/handling"
	"github.com/buger/jsonparser"
	"github.com/mymmrac/telego"
	"log/slog"
	"strconv"
	"strings"
	"time"
)

type RentUsecase interface {
	RentBook(ctx context.Context, input dto.RentInput) (rentService.RentBookOutput, error)
	FindBook(ctx context.Context, input dto.FindBookInput) (rentService.FindBookOutput, error)
	ConfirmRent(ctx context.Context, input dto.ConfirmRentInput) (rentService.ConfirmRentOutput, error)
	GetDebt(ctx context.Context, input dto.GetDebtInput) (rentService.GetDebtOutput, error)
	CheckRent(ctx context.Context, input dto.CheckRentInput) (rentService.CheckRentOutput, error)
	ConfirmReturn(ctx context.Context, input dto.ConfirmReturnInput) (rentService.ConfirmReturnOutput, error)
	MyRents(ctx context.Context, input dto.MyRentsInput) (rentService.MyRentsOutput, error)
}

type RentKeyboard interface {
	FindBook() *telego.InlineKeyboardMarkup
	Menu() *telego.InlineKeyboardMarkup
	Switch(offsetNext int, offsetPred int, id ...string) *telego.InlineKeyboardMarkup
	Success() *telego.InlineKeyboardMarkup
}

type RentHandler struct {
	builder          Builder
	router           Router
	question         Question
	callbackQuestion CallbackQuestion
	usecase          RentUsecase
	bookUsecase      BooksUsecase
	keyboard         RentKeyboard
}

func NewRentHandler(builder Builder, router Router, question Question, callback CallbackQuestion, rentUsecase RentUsecase, bookUsecase BooksUsecase, keyboard RentKeyboard) *RentHandler {
	return &RentHandler{
		builder:          builder,
		router:           router,
		question:         question,
		callbackQuestion: callback,
		usecase:          rentUsecase,
		bookUsecase:      bookUsecase,
		keyboard:         keyboard,
	}
}

func (h *RentHandler) AddGroup(handlerGroup handling.HandlersGroup) {
	h.router.AddGroup(handlerGroup)
}

func (h *RentHandler) Register() {
	regGroup := handling.NewHandlersGroup()
	rentBook := regGroup.NewHandler(h.RentBook)
	rentBook.WithCommand("/rent")
	rentBook.Question()
	myRents := regGroup.NewHandler(h.MyRents)
	myRents.WithCommand("/rented")
	nextBtn := regGroup.NewHandler(h.Next)
	nextBtn.WithCommand("/nextrents")
	predBtn := regGroup.NewHandler(h.Pred)
	predBtn.WithCommand("/predrents")
	cancelRent := regGroup.NewHandler(h.ConfirmReturn)
	cancelRent.WithCommand("/returnbook")
	cancelRent.Question()
	h.AddGroup(regGroup)
}

func (h *RentHandler) RentBook(ctx context.Context, msg telego.Update) {
	err := h.builder.NewCallbackMessage(msg.CallbackQuery, "")
	if err != nil {
		h.builder.NewMessage(msg, "Попробуйте заново.", h.keyboard.FindBook())
		slog.Error(err.Error())
		return
	}
	bookID, err := jsonparser.GetString([]byte(msg.CallbackQuery.Data), "id")
	if err != nil {
		h.builder.NewMessage(msg, "Попробуйте заново.", h.keyboard.FindBook())
		slog.Error(err.Error())
		return
	}
	bookIDInt, err := strconv.Atoi(bookID)
	if err != nil {
		h.builder.NewMessage(msg, "Попробуйте заново.", h.keyboard.FindBook())
		slog.Error(err.Error())
		return
	}
	inputCheck := dto.NewCheckRentInput(msg.CallbackQuery.From.ID, int64(bookIDInt))
	check, err := h.usecase.CheckRent(ctx, inputCheck)
	if err != nil || check.Status == 404 || check.ID != 0 {
		if check.ID != 0 {
			h.builder.NewMessage(msg, fmt.Sprintf("Вы уже создали запрос на аренду данной книги, если вы её до сих пор не получили, то подойдите к библиотекарю и продиктуйте номер: %d.", check.ID), h.keyboard.FindBook())
			return
		} else if check.ID == 0 && check.Error != "rpc error: code = Unknown desc = rent is not found" {
			h.builder.NewMessage(msg, "Попробуйте заново.", h.keyboard.FindBook())
			slog.Error(err.Error())
			return
		}
	}
	_, err = h.builder.NewMessage(msg, "Введите количество дней (от 1 до 30), через которые книга будет возвращена:", nil)
	if err != nil {
		h.builder.NewMessage(msg, "Попробуйте заново.", h.keyboard.FindBook())
		slog.Error(err.Error())
		return
	}
	answer, c := h.question.NewQuestion(msg)
	defer c()
	days, ok := <-answer
	if !ok || days.Text == "" {
		h.builder.NewMessage(msg, "Попробуйте заново.", h.keyboard.FindBook())
		return
	}
	daysInt, err := strconv.Atoi(days.Text)
	if err != nil {
		h.builder.NewMessage(msg, "Попробуйте ввести количество дней заново.", h.keyboard.FindBook())
		return
	}
	input := dto.NewRentInput(bookID, msg.CallbackQuery.From.ID, time.Now().Unix()+int64(daysInt*24*60*60))
	rent, err := h.usecase.RentBook(ctx, input)
	if err != nil || rent.Status == 404 || rent.ID == 0 {
		h.builder.NewMessage(msg, "Попробуйте заново.", h.keyboard.FindBook())
		slog.Error(err.Error())
		return
	}
	inputBook := dto.NewByInput(0, model.NewFindBook("", "", "", int64(bookIDInt)))
	book, err := h.bookUsecase.FindBy(ctx, inputBook)
	if err != nil {
		h.builder.NewMessage(msg, fmt.Sprintf("Запрос на аренду книги создан, но информации о книге не найдено.\n\nЧтобы подтвердить аренду книги подойдите в библиотеку и продиктуйте номер: %d", rent.ID), h.keyboard.FindBook())
		return
	}
	var findBooks []string
	var ids []string
	for _, book := range book.Book {
		ids = append(ids, strconv.Itoa(book.ID))
		findBooks = append(findBooks, fmt.Sprintf("ID: %d\nISBN: %s\nАвтор: %s\nНазвание: %s\nКоличество в библиотеке (шт): %d", book.ID, book.ISBN, book.Author, book.Name, book.Count))
	}
	h.builder.NewMessage(msg, fmt.Sprintf("Запрос на аренду книги создан. Информация о книге:\n\n%v\n\nЧтобы подтвердить аренду книги подойдите в библиотеку и продиктуйте номер: %d", strings.Join(findBooks, "\n\n"), rent.ID), h.keyboard.Menu())
}

func (h *RentHandler) GetDebt(ctx context.Context) {
	timeNow := time.Now().Unix()
	i := 0
	var debters []string
	var notify map[int64][]string
	notify = make(map[int64][]string)
	for i%5 == 0 {
		input := dto.NewGetDebtInput(timeNow, int64(i))
		debts, err := h.usecase.GetDebt(ctx, input)
		if err != nil {
			slog.Error(err.Error())
			break
		}
		if len(debts.Debt) > 0 {
			for _, debt := range debts.Debt {
				debters = append(debters, fmt.Sprintf("ISBN: %s,\nАвтор: %s,\nНазвание: %s,\nСсылка для перехода в диалог с арендатором: t.me/%s (после \"t.me/\" так же идёт его номер телефона),\nУчебный класс арендатора: %s;", debt.Books.ISBN, debt.Books.Author, debt.Books.Name, debt.Users.Phone, debt.Users.Class))
				notify[debt.Users.ID] = append(notify[debt.Users.ID], fmt.Sprintf("ISBN: %s,\nАвтор: %s,\nНазвание: %s.", debt.Books.ISBN, debt.Books.Author, debt.Books.Name))
			}
		}
		i += len(debts.Debt)
	}
	for k, v := range notify {
		h.builder.NewMessageByID(k, fmt.Sprintf("Ежедневное уведомление: вы должны вернуть книги(-у) со следующими данными в библиотеку, так как срок аренды истёк.\n\n%v", strings.Join(v, "\n\n")), nil)
	}
	if len(debters) == 0 {
		return
	}
	h.builder.NewMessageByID(1077777665, fmt.Sprintf("Обнаружены люди, у которых истёк срок аренды книги.\n\n%v", strings.Join(debters, "\n\n")), nil)
}

func (h *RentHandler) MyRents(ctx context.Context, msg telego.Update) {
	err := h.builder.NewCallbackMessage(msg.CallbackQuery, "")
	if err != nil {
		h.builder.NewMessage(msg, "Попробуйте заново.", h.keyboard.FindBook())
		slog.Error(err.Error())
		return
	}
	offset := 0
	input := dto.NewMyRentsInput(msg.CallbackQuery.From.ID, int64(offset))
	rents, err := h.usecase.MyRents(ctx, input)
	if len(rents.Rents) == 0 {
		h.builder.NewMessage(msg, "Вы не брали книг из библиотеки!", nil)
		return
	}
	var msgs []string
	var ids []string
	for _, rent := range rents.Rents {
		var isGet string
		if rent.IsGet == false {
			isGet = "нет"
		} else {
			isGet = "да"
		}
		normalTime := time.Unix(rent.ReturnAt, 0).UTC()
		msgs = append(msgs, fmt.Sprintf("Номер для подтверждения/отмены аренды: %d,\nISBN: %s,\nАвтор: %s,\nНазвание: %s,\nПолучена: %s,\nДата возврата: %v", rent.ID, rent.Books.ISBN, rent.Books.Author, rent.Books.Name, isGet, normalTime))
		if rent.IsGet == false {
			ids = append(ids, strconv.Itoa(int(rent.ID)))
		}
	}
	h.builder.NewMessage(msg, fmt.Sprintf("Вот данные о ваших книгах, которые вы арендовали (или создали запрос на их аренду):\n\n%v\n\nЕсли вы хотите отменить запрос на аренду книги, которую вы ещё не получили, то нажмите на кнопку с её ID.", strings.Join(msgs, "\n\n")), h.keyboard.Switch(offset+5, offset, ids...))
	return
}

func (h *RentHandler) Next(ctx context.Context, msg telego.Update) {
	err := h.builder.NewCallbackMessage(msg.CallbackQuery, "")
	if err != nil {
		h.builder.NewMessage(msg, "Попробуйте заново.", h.keyboard.FindBook())
		slog.Error(err.Error())
		return
	}

	offset, err := jsonparser.GetInt([]byte(msg.CallbackQuery.Data), "offset")
	if err != nil {
		h.builder.NewMessage(msg, "Попробуйте заново.", h.keyboard.FindBook())
		slog.Error(err.Error())
		return
	}

	input := dto.NewMyRentsInput(msg.CallbackQuery.From.ID, offset)
	rents, err := h.usecase.MyRents(ctx, input)
	var msgs []string
	var ids []string
	for _, rent := range rents.Rents {
		var isGet string
		if rent.IsGet == false {
			isGet = "нет"
		} else {
			isGet = "да"
		}
		normalTime := time.Unix(rent.ReturnAt, 0).UTC()
		msgs = append(msgs, fmt.Sprintf("Номер для подтверждения/отмены аренды: %d,\nISBN: %s,\nАвтор: %s,\nНазвание: %s,\nПолучена: %s,\nДата возврата: %v", rent.ID, rent.Books.ISBN, rent.Books.Author, rent.Books.Name, isGet, normalTime))
		if rent.IsGet == false {
			ids = append(ids, strconv.Itoa(int(rent.ID)))
		}
	}
	h.builder.NewMessage(msg, fmt.Sprintf("Вот данные о ваших книгах, которые вы арендовали (или создали запрос на их аренду):\n\n%v\n\nЕсли вы хотите отменить запрос на аренду книги, которую вы ещё не получили, то нажмите на кнопку с её ID.", strings.Join(msgs, "\n\n")), h.keyboard.Switch(int(offset+5), int(offset-5), ids...))
	return
}

func (h *RentHandler) Pred(ctx context.Context, msg telego.Update) {
	err := h.builder.NewCallbackMessage(msg.CallbackQuery, "")
	if err != nil {
		h.builder.NewMessage(msg, "Попробуйте заново.", h.keyboard.FindBook())
		slog.Error(err.Error())
		return
	}

	offset, err := jsonparser.GetInt([]byte(msg.CallbackQuery.Data), "offset")
	if err != nil {
		h.builder.NewMessage(msg, "Попробуйте заново.", h.keyboard.FindBook())
		slog.Error(err.Error())
		return
	}

	input := dto.NewMyRentsInput(msg.CallbackQuery.From.ID, offset)
	rents, err := h.usecase.MyRents(ctx, input)
	var msgs []string
	var ids []string
	for _, rent := range rents.Rents {
		var isGet string
		if rent.IsGet == false {
			isGet = "нет"
		} else {
			isGet = "да"
		}
		normalTime := time.Unix(rent.ReturnAt, 0).UTC()
		msgs = append(msgs, fmt.Sprintf("Номер для подтверждения/отмены аренды: %d,\nISBN: %s,\nАвтор: %s,\nНазвание: %s,\nПолучена: %s,\nДата возврата: %v", rent.ID, rent.Books.ISBN, rent.Books.Author, rent.Books.Name, isGet, normalTime))
		if rent.IsGet == false {
			ids = append(ids, strconv.Itoa(int(rent.ID)))
		}
	}
	var predOffset int64
	if offset > 0 {
		predOffset = offset - 5
	} else {
		predOffset = 0
	}
	h.builder.NewMessage(msg, fmt.Sprintf("Вот данные о ваших книгах, которые вы арендовали (или создали запрос на их аренду):\n\n%v\n\nЕсли вы хотите отменить запрос на аренду книги, которую вы ещё не получили, то нажмите на кнопку с её ID.", strings.Join(msgs, "\n\n")), h.keyboard.Switch(int(offset+5), int(predOffset), ids...))
	return
}

func (h *RentHandler) ConfirmReturn(ctx context.Context, msg telego.Update) {
	err := h.builder.NewCallbackMessage(msg.CallbackQuery, "")
	if err != nil {
		h.builder.NewMessage(msg, "Попробуйте заново.", h.keyboard.FindBook())
		slog.Error(err.Error())
		return
	}

	id, err := jsonparser.GetString([]byte(msg.CallbackQuery.Data), "id")
	if err != nil {
		h.builder.NewMessage(msg, "Попробуйте заново.", h.keyboard.FindBook())
		slog.Error(err.Error())
		return
	}
	idInt, err := strconv.Atoi(id)
	if err != nil {
		slog.Error(err.Error())
		h.builder.NewMessage(msg, "Попробуйте заново.", h.keyboard.FindBook())
		return
	}
	input := dto.NewFindBookInput(int64(idInt))
	book, err := h.usecase.FindBook(ctx, input)
	if err != nil {
		slog.Error(err.Error())
		h.builder.NewMessage(msg, "Попробуйте заново.", h.keyboard.FindBook())
		return
	}
	_, err = h.builder.NewMessage(msg, fmt.Sprintf("Вы точно хотите подтвердить отмену запроса на аренду книги? Вот информация о ней:\n\nID: %d\nISBN: %s\nАвтор: %s\nНазвание: %s\nКоличество в библиотеке (шт): %d", book.Model[0].Books.ID, book.Model[0].Books.ISBN, book.Model[0].Books.Author, book.Model[0].Books.Name, book.Model[0].Books.Count), h.keyboard.Success())
	if err != nil {
		slog.Error(err.Error())
		return
	}
	callbackAnswers, cl := h.callbackQuestion.NewQuestion(msg)
	defer cl()
	answer, ok := <-callbackAnswers
	if !ok {
		h.builder.NewMessage(msg, "Попробуйте заново.", h.keyboard.FindBook())
		return
	}
	if answer.CallbackQuery.Data == "{\"command\":\"/accept\"}" {
		err = h.builder.NewCallbackMessage(answer.CallbackQuery, "")
		if err != nil {
			slog.Error(err.Error())
			return
		}
		input := dto.NewConfirmReturnInput(int64(idInt))
		res, err := h.usecase.ConfirmReturn(ctx, input)
		if err != nil || res.Status == 404 {
			h.builder.NewMessage(msg, "Попробуйте заново позже.", h.keyboard.FindBook())
			slog.Error(err.Error())
			return
		}
		h.builder.NewMessage(msg, "Вы успешно подтвердили отмену запроса на аренду книги!", h.keyboard.Menu())
	} else {
		h.builder.NewMessage(msg, "Вот ваш функционал:", h.keyboard.Menu())
	}
	return
}
