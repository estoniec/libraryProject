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
}

type RentKeyboard interface {
	FindBook() *telego.InlineKeyboardMarkup
	Menu() *telego.InlineKeyboardMarkup
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
	input := dto.NewGetDebtInput(timeNow)
	debts, err := h.usecase.GetDebt(ctx, input)
	if err != nil {
		slog.Error(err.Error())
		return
	}
	if len(debts.Debt) > 0 {
		var debters []string
		for _, debt := range debts.Debt {
			debters = append(debters, fmt.Sprintf("ISBN: %s,\nАвтор: %s,\nНазвание: %s,\nСсылка для перехода в диалог с арендатором: t.me/%s (после \"t.me/\" так же идёт его номер телефона),\nУчебный класс арендатора: %s;", debt.Books.ISBN, debt.Books.Author, debt.Books.Name, debt.Users.Phone, debt.Users.Class))
			h.builder.NewMessageByID(debt.Users.ID, fmt.Sprintf("Ежедневное уведомление: вы должны вернуть книгу со следующими данными в библиотеку, так как срок аренды истёк.\n\nISBN: %s,\nАвтор: %s,\nНазвание: %s.", debt.Books.ISBN, debt.Books.Author, debt.Books.Name), nil)
		}
		h.builder.NewMessageByID(1077777665, fmt.Sprintf("Обнаружены люди, у которых истёк срок аренды книги.\n\n%v", strings.Join(debters, "\n\n")), nil)
	}
}

func (h *RentHandler) MyRents(ctx context.Context, msg telego.Update) {
	err := h.builder.NewCallbackMessage(msg.CallbackQuery, "")
	if err != nil {
		h.builder.NewMessage(msg, "Попробуйте заново.", h.keyboard.FindBook())
		slog.Error(err.Error())
		return
	}
}

// TODO сделать отмену запроса аренды книги
