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
	keyboard         AdminKeyboard
}

func NewAdminHandler(builder Builder, router Router, question Question, callback CallbackQuestion, regUsecase RegUsecase, bookUsecase BooksUsecase, keyboard AdminKeyboard) *AdminHandler {
	return &AdminHandler{
		builder:          builder,
		router:           router,
		question:         question,
		callbackQuestion: callback,
		regUsecase:       regUsecase,
		bookUsecase:      bookUsecase,
		keyboard:         keyboard,
	}
}

func (h *AdminHandler) AddGroup(handlerGroup handling.HandlersGroup) {
	h.router.AddGroup(handlerGroup)
}

func (h *AdminHandler) Register() {
	regGroup := handling.NewHandlersGroup()
	regGroup.NewHandler(h.AddBook, "Добавить книгу")
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
	_, err = h.builder.NewMessage(msg, "Вы перешли в создание книги. Для начала введите её ISBN:", nil)
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

// TODO додлеть EditCountBook, DeleteBook
