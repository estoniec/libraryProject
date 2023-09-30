package v1

import (
	"context"
	"fmt"
	dto2 "gateway/internal/domain/books/dto"
	"gateway/pkg/adapters/handling"
	"github.com/mymmrac/telego"
	"log/slog"
	"strconv"
	"strings"
)

type BooksService interface {
	FindByISBN(ctx context.Context, input dto2.FindByISBNInput) (dto2.FindByISBNOutput, error)
	FindByAuthor(ctx context.Context, input dto2.FindByAuthorInput) (dto2.FindByAuthorOutput, error)
	FindByName(ctx context.Context, input dto2.FindByNameInput) (dto2.FindByNameOutput, error)
}

type BooksKeyboard interface {
	FindBook() *telego.InlineKeyboardMarkup
	FindBy(fromID int64, findType string, something string, id ...string) *telego.InlineKeyboardMarkup
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
	findByAuthor := regGroup.NewHandler(h.FindByAuthor)
	findByAuthor.WithCommand("/findbyauthor")
	h.AddGroup(regGroup)
}

func (h *BooksHandler) AvailableBooks(ctx context.Context, msg telego.Update) {
	err := h.builder.NewCallbackMessage(msg.CallbackQuery, "")
	if err != nil {
		slog.Error(err.Error())
		return
	}
	_, err = h.builder.NewMessage(msg, "Сначала нужно найти книгу. Выберите тип поиска:", h.keyboard.FindBook())
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
	h.builder.NewMessage(msg, fmt.Sprintf("Книга найдена, вот информация о ней:\n\nID: %d,\nISBN: %s,\nАвтор: %s,\nНазвание: %s,\nКоличество в библиотеке (шт): %d.", book.Book.ID, book.Book.ISBN, book.Book.Author, book.Book.Name, book.Book.Count), h.keyboard.FindBook())
	return
}

func (h *BooksHandler) FindByAuthor(ctx context.Context, msg telego.Update) {
	err := h.builder.NewCallbackMessage(msg.CallbackQuery, "")
	if err != nil {
		h.builder.NewMessage(msg, "Попробуйте заново.", h.keyboard.FindBook())
		slog.Error(err.Error())
		return
	}
	_, err = h.builder.NewMessage(msg, "Введите автора книги (например, \"А. А. Блок\"):", nil)
	if err != nil {
		h.builder.NewMessage(msg, "Попробуйте заново.", h.keyboard.FindBook())
		slog.Error(err.Error())
		return
	}
	answer, c := h.question.NewQuestion(msg)
	defer c()
	author, ok := <-answer
	if !ok || author.Text == "" {
		h.builder.NewMessage(msg, "Попробуйте заново.", h.keyboard.FindBook())
		return
	}
	dto := dto2.NewAuthorInput(author.Text, 0)
	books, err := h.service.FindByAuthor(ctx, dto)
	if err != nil || books.Status != 200 {
		if err.Error() == "rpc error: code = Unknown desc = book is not found" {
			h.builder.NewMessage(msg, "Книг от такого автора не существует.", h.keyboard.FindBook())
			slog.Error(err.Error())
			return
		}
		h.builder.NewMessage(msg, "Попробуйте заново.", h.keyboard.FindBook())
		slog.Error(err.Error())
		return
	}
	var findBooks []string
	var ids []string
	for _, book := range books.Book {
		ids = append(ids, strconv.Itoa(book.ID))
		findBooks = append(findBooks, fmt.Sprintf("ID: %d,\nISBN: %s,\nАвтор: %s,\nНазвание: %s,\nКоличество в библиотеке (шт): %d", book.ID, book.ISBN, book.Author, book.Name, book.Count))
	}
	h.builder.NewMessage(msg, fmt.Sprintf("Книги найдены, вот информация о них:\n\n%v", strings.Join(findBooks, ";\n\n")), h.keyboard.FindBy(msg.CallbackQuery.From.ID, "author", author.Text, ids...))
	return
}

func (h *BooksHandler) FindByName(ctx context.Context, msg telego.Update) {
	err := h.builder.NewCallbackMessage(msg.CallbackQuery, "")
	if err != nil {
		h.builder.NewMessage(msg, "Попробуйте заново.", h.keyboard.FindBook())
		slog.Error(err.Error())
		return
	}
	_, err = h.builder.NewMessage(msg, "Введите название книги (например, \"Россия\"):", nil)
	if err != nil {
		h.builder.NewMessage(msg, "Попробуйте заново.", h.keyboard.FindBook())
		slog.Error(err.Error())
		return
	}
	answer, c := h.question.NewQuestion(msg)
	defer c()
	name, ok := <-answer
	if !ok || name.Text == "" {
		h.builder.NewMessage(msg, "Попробуйте заново.", h.keyboard.FindBook())
		return
	}
	dto := dto2.NewNameInput(name.Text, 0)
	books, err := h.service.FindByName(ctx, dto)
	if err != nil || books.Status != 200 {
		if err.Error() == "rpc error: code = Unknown desc = book is not found" {
			h.builder.NewMessage(msg, "Книг с таким названием не существует.", h.keyboard.FindBook())
			slog.Error(err.Error())
			return
		}
		h.builder.NewMessage(msg, "Попробуйте заново.", h.keyboard.FindBook())
		slog.Error(err.Error())
		return
	}
	var findBooks []string
	var ids []string
	for _, book := range books.Book {
		ids = append(ids, strconv.Itoa(book.ID))
		findBooks = append(findBooks, fmt.Sprintf("ID: %d,\nISBN: %s,\nАвтор: %s,\nНазвание: %s,\nКоличество в библиотеке (шт): %d", book.ID, book.ISBN, book.Author, book.Name, book.Count))
	}
	h.builder.NewMessage(msg, fmt.Sprintf("Книги найдены, вот информация о них:\n\n%v", strings.Join(findBooks, ";\n\n")), h.keyboard.FindBy(msg.CallbackQuery.From.ID, "name", name.Text, ids...))
	return
}
