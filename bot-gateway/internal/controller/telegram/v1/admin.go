package v1

import (
	"context"
	"fmt"
	"gateway/internal/controller/telegram/dto"
	"gateway/internal/domain/books/model"
	"gateway/pkg/adapters/handling"
	"github.com/mymmrac/telego"
	"log/slog"
	"strconv"
)

type AdminKeyboard interface {
	Admin() *telego.ReplyKeyboardMarkup
	SuccessAdd() *telego.InlineKeyboardMarkup
}

type AdminHandler struct {
	builder          Builder
	router           Router
	question         Question
	callbackQuestion CallbackQuestion
	regUsecase       RegUsecase
	bookUsecase      BooksUsecase
	rentUsecase      RentUsecase
	keyboard         AdminKeyboard
}

func NewAdminHandler(builder Builder, router Router, question Question, callback CallbackQuestion, regUsecase RegUsecase, bookUsecase BooksUsecase, rentUsecase RentUsecase, keyboard AdminKeyboard) *AdminHandler {
	return &AdminHandler{
		builder:          builder,
		router:           router,
		question:         question,
		callbackQuestion: callback,
		regUsecase:       regUsecase,
		bookUsecase:      bookUsecase,
		rentUsecase:      rentUsecase,
		keyboard:         keyboard,
	}
}

func (h *AdminHandler) AddGroup(handlerGroup handling.HandlersGroup) {
	h.router.AddGroup(handlerGroup)
}

func (h *AdminHandler) Register() {
	regGroup := handling.NewHandlersGroup()
	regGroup.NewHandler(h.AddBook, "Добавить книгу")
	regGroup.NewHandler(h.EditCountBook, "Изменить количество книг")
	regGroup.NewHandler(h.DeleteBook, "Удалить книгу")
	regGroup.NewHandler(h.ConfirmRent, "Подтвердить аренду книги")
	keyboard := regGroup.NewHandler(h.GetKeyboard, "/getkeyboard")
	keyboard.WithCommand("/cancel")
	h.AddGroup(regGroup)
}

func (h *AdminHandler) GetKeyboard(ctx context.Context, msg telego.Update) {
	var fromID int64
	if msg.CallbackQuery != nil {
		err := h.builder.NewCallbackMessage(msg.CallbackQuery, "")
		if err != nil {
			slog.Error(err.Error())
			return
		}
		fromID = msg.CallbackQuery.From.ID
	} else {
		fromID = msg.Message.From.ID
	}
	dtoCheck := dto.NewCheckRoleInput(fromID)
	role, err := h.regUsecase.CheckRole(ctx, dtoCheck)
	if err != nil || role.Status == 404 {
		slog.Error(role.Error)
		return
	}
	if role.Role < 2 {
		_, err = h.builder.NewMessage(msg, "Для использования данной команды у вас недостаточно прав", nil)
		if err != nil {
			slog.Error(err.Error())
			return
		}
		return
	}
	_, err = h.builder.NewMessageWithKeyboard(msg, "Вот ваш функционал:", h.keyboard.Admin())
	if err != nil {
		slog.Error(err.Error())
		return
	}
}

func (h *AdminHandler) AddBook(ctx context.Context, msg telego.Update) {
	dtoCheck := dto.NewCheckRoleInput(msg.Message.From.ID)
	role, err := h.regUsecase.CheckRole(ctx, dtoCheck)
	if err != nil || role.Status == 404 {
		slog.Error(role.Error)
		return
	}
	if role.Role < 2 {
		_, err = h.builder.NewMessage(msg, "Для использования данной команды у вас недостаточно прав", nil)
		if err != nil {
			slog.Error(err.Error())
			return
		}
		return
	}
	_, err = h.builder.NewMessage(msg, "Вы перешли в создание книги. Для начала введите её ISBN (пример ввода: \"5080020229\"), который указан на 1 или 2 странице книги:", nil)
	if err != nil {
		slog.Error(err.Error())
		return
	}
	answers, c := h.question.NewQuestion(msg)
	defer c()
	isbn, ok := <-answers
	if !ok || isbn.Text == "" {
		h.builder.NewMessage(msg, "Попробуйте ввести ISBN заново.", nil)
		return
	}
	_, err = h.builder.NewMessage(msg, "Теперь введите автора книги (пример: \"А. А. Блок\")", nil)
	if err != nil {
		slog.Error(err.Error())
		return
	}
	author, ok := <-answers
	if !ok || author.Text == "" {
		h.builder.NewMessage(msg, "Попробуйте ввести автора заново.", nil)
		return
	}
	_, err = h.builder.NewMessage(msg, "Теперь введите название книги (пример: \"Россия\")", nil)
	if err != nil {
		slog.Error(err.Error())
		return
	}
	name, ok := <-answers
	if !ok || name.Text == "" {
		h.builder.NewMessage(msg, "Попробуйте ввести название заново.", nil)
		return
	}
	_, err = h.builder.NewMessage(msg, "Последний шаг! Введите количество книг в наличии:", nil)
	if err != nil {
		slog.Error(err.Error())
		return
	}
	count, ok := <-answers
	if !ok || count.Text == "" {
		h.builder.NewMessage(msg, "Попробуйте ввести количество книг заново.", nil)
		return
	}
	countInt, err := strconv.Atoi(count.Text)
	if err != nil {
		h.builder.NewMessage(msg, "Попробуйте ввести количество книг заново.", nil)
		return
	}
	_, err = h.builder.NewMessage(msg, fmt.Sprintf("Вы хотите добавить в базу данных следующую книгу:\nISBN: %s\nАвтор: %s\nНазвание: %s\nКоличество (шт.): %s", isbn.Text, author.Text, name.Text, count.Text), h.keyboard.SuccessAdd())
	if err != nil {
		slog.Error(err.Error())
		return
	}
	callbackAnswers, cl := h.callbackQuestion.NewQuestion(msg)
	defer cl()
	answer, ok := <-callbackAnswers
	if !ok {
		h.builder.NewMessage(msg, "Попробуйте заново.", nil)
		return
	}
	if answer.CallbackQuery.Data == "{\"command\":\"/accept\"}" {
		err = h.builder.NewCallbackMessage(answer.CallbackQuery, "")
		if err != nil {
			slog.Error(err.Error())
			return
		}
		input := dto.NewAddBookInput(model.NewBook(isbn.Text, countInt, name.Text, author.Text))
		res, err := h.bookUsecase.AddBook(ctx, input)
		if err != nil || res.Status == 404 {
			h.builder.NewMessage(msg, "Попробуйте заново позже.", nil)
			slog.Error(err.Error())
			return
		}
		h.builder.NewMessageWithKeyboard(msg, "Вы успешно добавили книгу в базу данных!", h.keyboard.Admin())
	} else {
		h.GetKeyboard(ctx, answer)
	}
	return
}

func (h *AdminHandler) EditCountBook(ctx context.Context, msg telego.Update) {
	dtoCheck := dto.NewCheckRoleInput(msg.Message.From.ID)
	role, err := h.regUsecase.CheckRole(ctx, dtoCheck)
	if err != nil || role.Status == 404 {
		slog.Error(role.Error)
		return
	}
	if role.Role < 2 {
		_, err = h.builder.NewMessage(msg, "Для использования данной команды у вас недостаточно прав", nil)
		if err != nil {
			slog.Error(err.Error())
			return
		}
		return
	}
	_, err = h.builder.NewMessage(msg, "Введите ISBN книги (пример ввода: \"5080020229\"), который указан на 1 или 2 странице книги:", nil)
	if err != nil {
		slog.Error(err.Error())
		return
	}
	answers, c := h.question.NewQuestion(msg)
	defer c()
	isbn, ok := <-answers
	if !ok || isbn.Text == "" {
		h.builder.NewMessage(msg, "Попробуйте ввести ISBN заново.", nil)
		return
	}
	_, err = h.builder.NewMessage(msg, "Теперь введите количество книг, которое вы хотите установить:", nil)
	if err != nil {
		slog.Error(err.Error())
		return
	}
	count, ok := <-answers
	if !ok || count.Text == "" {
		h.builder.NewMessage(msg, "Попробуйте ввести автора заново.", nil)
		return
	}
	countInt, err := strconv.Atoi(count.Text)
	if err != nil {
		h.builder.NewMessage(msg, "Попробуйте ввести количество книг заново.", nil)
		return
	}
	_, err = h.builder.NewMessage(msg, fmt.Sprintf("Вы точно хотите изменить количество книг у книги с ISBN: %s на %s (шт.)?", isbn.Text, count.Text), h.keyboard.SuccessAdd())
	if err != nil {
		slog.Error(err.Error())
		return
	}
	callbackAnswers, cl := h.callbackQuestion.NewQuestion(msg)
	defer cl()
	answer, ok := <-callbackAnswers
	if !ok {
		h.builder.NewMessage(msg, "Попробуйте заново.", nil)
		return
	}
	if answer.CallbackQuery.Data == "{\"command\":\"/accept\"}" {
		err = h.builder.NewCallbackMessage(answer.CallbackQuery, "")
		if err != nil {
			slog.Error(err.Error())
			return
		}
		input := dto.NewEditCountBookInput(isbn.Text, countInt)
		res, err := h.bookUsecase.EditCountBook(ctx, input)
		if err != nil || res.Status == 404 {
			h.builder.NewMessage(msg, "Попробуйте заново позже.", nil)
			slog.Error(err.Error())
			return
		}
		h.builder.NewMessageWithKeyboard(msg, "Вы успешно изменили количество книг!", h.keyboard.Admin())
	} else {
		h.GetKeyboard(ctx, answer)
	}
	return
}

func (h *AdminHandler) DeleteBook(ctx context.Context, msg telego.Update) {
	dtoCheck := dto.NewCheckRoleInput(msg.Message.From.ID)
	role, err := h.regUsecase.CheckRole(ctx, dtoCheck)
	if err != nil || role.Status == 404 {
		slog.Error(role.Error)
		return
	}
	if role.Role < 2 {
		_, err = h.builder.NewMessage(msg, "Для использования данной команды у вас недостаточно прав", nil)
		if err != nil {
			slog.Error(err.Error())
			return
		}
		return
	}
	_, err = h.builder.NewMessage(msg, "Введите ISBN книги (пример ввода: \"5080020229\"), который указан на 1 или 2 странице книги:", nil)
	if err != nil {
		slog.Error(err.Error())
		return
	}
	answers, c := h.question.NewQuestion(msg)
	defer c()
	isbn, ok := <-answers
	if !ok || isbn.Text == "" {
		h.builder.NewMessage(msg, "Попробуйте ввести ISBN заново.", nil)
		return
	}
	_, err = h.builder.NewMessage(msg, fmt.Sprintf("Вы точно хотите удалить книгу с ISBN: %s?", isbn.Text), h.keyboard.SuccessAdd())
	if err != nil {
		slog.Error(err.Error())
		return
	}
	callbackAnswers, cl := h.callbackQuestion.NewQuestion(msg)
	defer cl()
	answer, ok := <-callbackAnswers
	if !ok {
		h.builder.NewMessage(msg, "Попробуйте заново.", nil)
		return
	}
	if answer.CallbackQuery.Data == "{\"command\":\"/accept\"}" {
		err = h.builder.NewCallbackMessage(answer.CallbackQuery, "")
		if err != nil {
			slog.Error(err.Error())
			return
		}
		input := dto.NewDeleteBookInput(isbn.Text)
		res, err := h.bookUsecase.DeleteBook(ctx, input)
		if err != nil || res.Status == 404 {
			h.builder.NewMessage(msg, "Попробуйте заново позже.", nil)
			slog.Error(err.Error())
			return
		}
		h.builder.NewMessageWithKeyboard(msg, "Вы успешно изменили количество книг!", h.keyboard.Admin())
	} else {
		h.GetKeyboard(ctx, answer)
	}
	return
}

func (h *AdminHandler) ConfirmRent(ctx context.Context, msg telego.Update) {
	dtoCheck := dto.NewCheckRoleInput(msg.Message.From.ID)
	role, err := h.regUsecase.CheckRole(ctx, dtoCheck)
	if err != nil || role.Status == 404 {
		slog.Error(role.Error)
		return
	}
	if role.Role < 2 {
		_, err = h.builder.NewMessage(msg, "Для использования данной команды у вас недостаточно прав", nil)
		if err != nil {
			slog.Error(err.Error())
			return
		}
		return
	}
	_, err = h.builder.NewMessage(msg, "Введите номер, который вам продиктует арендатор:", nil)
	if err != nil {
		slog.Error(err.Error())
		return
	}
	answers, c := h.question.NewQuestion(msg)
	defer c()
	id, ok := <-answers
	if !ok || id.Text == "" {
		h.builder.NewMessage(msg, "Попробуйте ввести номер заново.", nil)
		return
	}
	idInt, err := strconv.Atoi(id.Text)
	if err != nil {
		h.builder.NewMessage(msg, "Попробуйте ввести номер заново.", nil)
		return
	}
	if err != nil {
		slog.Error(err.Error())
		return
	}
	input := dto.NewFindBookInput(int64(idInt))
	book, err := h.rentUsecase.FindBook(ctx, input)
	if err != nil {
		slog.Error(err.Error())
		h.builder.NewMessage(msg, "Попробуйте заново.", nil)
		return
	}
	_, err = h.builder.NewMessage(msg, fmt.Sprintf("Вы точно хотите подтвердить аренду книги? Вот информация о ней:\n\nID: %d\nISBN: %s\nАвтор: %s\nНазвание: %s\nКоличество в библиотеке (шт): %d", book.Book.GetID(), book.Book.GetISBN(), book.Book.GetAuthor(), book.Book.GetName(), book.Book.GetCount()), h.keyboard.SuccessAdd())
	if err != nil {
		slog.Error(err.Error())
		return
	}
	callbackAnswers, cl := h.callbackQuestion.NewQuestion(msg)
	defer cl()
	answer, ok := <-callbackAnswers
	if !ok {
		h.builder.NewMessage(msg, "Попробуйте заново.", nil)
		return
	}
	if answer.CallbackQuery.Data == "{\"command\":\"/accept\"}" {
		err = h.builder.NewCallbackMessage(answer.CallbackQuery, "")
		if err != nil {
			slog.Error(err.Error())
			return
		}
		input := dto.NewConfirmRentInput(int64(idInt))
		res, err := h.rentUsecase.ConfirmRent(ctx, input)
		if err != nil || res.Status == 404 {
			h.builder.NewMessage(msg, "Попробуйте заново позже.", nil)
			slog.Error(err.Error())
			return
		}
		inputEdit := dto.NewEditCountBookInput(book.Book.GetISBN(), int(book.Book.GetCount())-1)
		resEdit, err := h.bookUsecase.EditCountBook(ctx, inputEdit)
		if err != nil || resEdit.Status == 404 {
			h.builder.NewMessage(msg, "Попробуйте заново позже.", nil)
			slog.Error(err.Error())
			return
		}
		h.builder.NewMessageWithKeyboard(msg, "Вы успешно подтвердили аренду книги!", h.keyboard.Admin())
	} else {
		h.GetKeyboard(ctx, answer)
	}
	return
}
