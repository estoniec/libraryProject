package v1

import (
	"context"
	"fmt"
	dto2 "gateway/internal/domain/books/dto"
	"gateway/pkg/adapters/handling"
	"github.com/mymmrac/telego"
	"log/slog"
)

type BooksService interface {
	FindByISBN(ctx context.Context, input dto2.FindByISBNInput) (dto2.FindByISBNOutput, error)
}

type BooksKeyboard interface {
	FindBook() *telego.InlineKeyboardMarkup
}

type BooksHandler struct {
	builder          Builder
	router           Router
	question         Question
	callbackQuestion CallbackQuestion
	service          BooksService
	keyboard         BooksKeyboard
}

func NewBooksHandler(builder Builder, router Router, question Question, callbackQuestion CallbackQuestion, service BooksService, keyboard BooksKeyboard) *BooksHandler {
	return &BooksHandler{
		builder:          builder,
		router:           router,
		question:         question,
		callbackQuestion: callbackQuestion,
		service:          service,
		keyboard:         keyboard,
	}
}

func (h *BooksHandler) AddGroup(handlerGroup handling.HandlersGroup) {
	h.router.AddGroup(handlerGroup)
}

func (h *BooksHandler) Register() {
	regGroup := handling.NewHandlersGroup()
	av := regGroup.NewHandler(h.AvailableBooks)
	av.WithCommand("/available")
	findByISBN := regGroup.NewHandler(h.FindByISBN)
	findByISBN.WithCommand("/findbyisbn")
	h.AddGroup(regGroup)
}

func (h *BooksHandler) AvailableBooks(ctx context.Context, msg telego.Update) {
	err := h.builder.NewCallbackMessage(msg.CallbackQuery, "")
	if err != nil {
		slog.Error(err.Error())
		return
	}
	_, err = h.builder.NewMessage(msg, "Выберите тип поиска:", h.keyboard.FindBook())
	if err != nil {
		slog.Error(err.Error())
		return
	}
	answer, c := h.callbackQuestion.NewQuestion(msg)
	defer c()
	_, ok := <-answer
	if !ok {
		h.builder.NewMessage(msg, "Повторите попытку позже.", h.keyboard.FindBook())
		return
	}
}

func (h *BooksHandler) FindByISBN(ctx context.Context, msg telego.Update) {
	err := h.builder.NewCallbackMessage(msg.CallbackQuery, "")
	if err != nil {
		h.builder.NewMessage(msg, "Попробуйте заново.", h.keyboard.FindBook())
		slog.Error(err.Error())
		return
	}
	_, err = h.builder.NewMessage(msg, "Введите ISBN книги:", nil)
	if err != nil {
		h.builder.NewMessage(msg, "Попробуйте заново.", h.keyboard.FindBook())
		slog.Error(err.Error())
		return
	}
	answer, c := h.question.NewQuestion(msg)
	defer c()
	isbn, ok := <-answer
	if !ok || isbn.Text == "" {
		h.builder.NewMessage(msg, "Попробуйте заново.", h.keyboard.FindBook())
		return
	}
	dto := dto2.NewISBNInput(isbn.Text)
	book, err := h.service.FindByISBN(ctx, dto)
	if err != nil || book.Status != 200 {
		if err.Error() == "rpc error: code = Unknown desc = book is not found" {
			h.builder.NewMessage(msg, "Книги с таким ISBN не существует.", h.keyboard.FindBook())
			slog.Error(err.Error())
			return
		}
		h.builder.NewMessage(msg, "Попробуйте заново.", h.keyboard.FindBook())
		slog.Error(err.Error())
		return
	}
	h.builder.NewMessage(msg, fmt.Sprintf("Книга найдена, вот информация о ней:\n\nISBN: %s,\nАвтор: %s,\nНазвание: %s,\nКоличество в библиотеке (шт): %d.", book.Book.ISBN, book.Book.Author, book.Book.Name, book.Book.Count), h.keyboard.FindBook())
	return
}
