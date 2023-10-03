package v1

import (
	"context"
	"fmt"
	dto2 "gateway/internal/domain/books/dto"
	"gateway/pkg/adapters/handling"
	"github.com/buger/jsonparser"
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
	FindBy(findType string, offsetNext int, offsetPred int, something string, id ...string) *telego.InlineKeyboardMarkup
	Find(id string) *telego.InlineKeyboardMarkup
	Menu() *telego.InlineKeyboardMarkup
}

type BooksHandler struct {
	builder          Builder
	router           Router
	question         Question
	callbackQuestion CallbackQuestion
	service          BooksService
	keyboard         BooksKeyboard
	Fetches          map[int64]int64
}

func NewBooksHandler(builder Builder, router Router, question Question, callbackQuestion CallbackQuestion, service BooksService, keyboard BooksKeyboard) *BooksHandler {
	return &BooksHandler{
		builder:          builder,
		router:           router,
		question:         question,
		callbackQuestion: callbackQuestion,
		service:          service,
		keyboard:         keyboard,
		Fetches:          make(map[int64]int64),
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
	nextBtn := regGroup.NewHandler(h.Next)
	nextBtn.WithCommand("/next")
	menu := regGroup.NewHandler(h.Menu)
	menu.WithCommand("/menu")
	h.AddGroup(regGroup)
}

func (h *BooksHandler) Menu(ctx context.Context, msg telego.Update) {
	if _, ok := h.Fetches[msg.CallbackQuery.From.ID]; ok {
		delete(h.Fetches, msg.CallbackQuery.From.ID)
		h.Fetches[msg.CallbackQuery.From.ID] = 0
	}
	err := h.builder.NewCallbackMessage(msg.CallbackQuery, "")
	if err != nil {
		slog.Error(err.Error())
		return
	}
	_, err = h.builder.NewMessage(msg, "Вот ваш функционал:", h.keyboard.Menu())
	if err != nil {
		slog.Error(err.Error())
		return
	}
	return
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
	h.builder.NewMessage(msg, fmt.Sprintf("Книга найдена, вот информация о ней:\n\nID: %d,\nISBN: %s,\nАвтор: %s,\nНазвание: %s,\nКоличество в библиотеке (шт): %d.", book.Book.ID, book.Book.ISBN, book.Book.Author, book.Book.Name, book.Book.Count), h.keyboard.Find(strconv.Itoa(book.Book.ID)))
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
	if _, ok := h.Fetches[msg.CallbackQuery.From.ID]; ok {
		delete(h.Fetches, msg.CallbackQuery.From.ID)
		h.Fetches[msg.CallbackQuery.From.ID] = 0
	} else {
		h.Fetches[msg.CallbackQuery.From.ID] = 0
	}
	dto := dto2.NewAuthorInput(author.Text, h.Fetches[msg.CallbackQuery.From.ID])
	books, err := h.service.FindByAuthor(ctx, dto)
	if err != nil || books.Status != 200 {
		if err.Error() == "rpc error: code = Unknown desc = book is not found" {
			h.builder.NewMessage(msg, "Книг от такого автора не существует.", h.keyboard.FindBook())
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
		findBooks = append(findBooks, fmt.Sprintf("ID: %d\nISBN: %s\nАвтор: %s\nНазвание: %s\nКоличество в библиотеке (шт): %d", book.ID, book.ISBN, book.Author, book.Name, book.Count))
	}
	h.builder.NewMessage(msg, fmt.Sprintf("Книги найдены, вот информация о них:\n\n%v", strings.Join(findBooks, "\n\n")), h.keyboard.FindBy("author", int(h.Fetches[msg.CallbackQuery.From.ID]), int(h.Fetches[msg.CallbackQuery.From.ID]), author.Text, ids...))
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
	if _, ok := h.Fetches[msg.CallbackQuery.From.ID]; ok {
		delete(h.Fetches, msg.CallbackQuery.From.ID)
		h.Fetches[msg.CallbackQuery.From.ID] = 0
	} else {
		h.Fetches[msg.CallbackQuery.From.ID] = 0
	}
	dto := dto2.NewNameInput(name.Text, h.Fetches[msg.CallbackQuery.From.ID])
	books, err := h.service.FindByName(ctx, dto)
	if err != nil || books.Status != 200 {
		if err.Error() == "rpc error: code = Unknown desc = book is not found" {
			h.builder.NewMessage(msg, "Книг с таким названием не существует.", h.keyboard.FindBook())
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
		findBooks = append(findBooks, fmt.Sprintf("ID: %d\nISBN: %s\nАвтор: %s\nНазвание: %s\nКоличество в библиотеке (шт): %d", book.ID, book.ISBN, book.Author, book.Name, book.Count))
	}
	h.builder.NewMessage(msg, fmt.Sprintf("Книги найдены, вот информация о них:\n\n%v\n\nЧтобы арендовать какую-то из этих книг, нажмите на кнопку с её ID.", strings.Join(findBooks, "\n\n")), h.keyboard.FindBy("name", int(h.Fetches[msg.CallbackQuery.From.ID]), int(h.Fetches[msg.CallbackQuery.From.ID]), name.Text, ids...))
	return
}

func (h *BooksHandler) RentBook(ctx context.Context, msg telego.Update) {
	err := h.builder.NewCallbackMessage(msg.CallbackQuery, "")
	if err != nil {
		h.builder.NewMessage(msg, "Попробуйте заново.", h.keyboard.FindBook())
		slog.Error(err.Error())
		return
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

	//dto := dto2.NewNameInput(days.Text, 0)
	//books, err := h.service.FindByName(ctx, dto)
	//if err != nil || books.Status != 200 {
	//	if err.Error() == "rpc error: code = Unknown desc = book is not found" {
	//		h.builder.NewMessage(msg, "Книг с таким названием не существует.", h.keyboard.FindBook())
	//		slog.Error(err.Error())
	//		return
	//	}
	//	h.builder.NewMessage(msg, "Попробуйте заново.", h.keyboard.FindBook())
	//	slog.Error(err.Error())
	//	return
	//}
	//var findBooks []string
	//var ids []string
	//for _, book := range books.Book {
	//	ids = append(ids, strconv.Itoa(book.ID))
	//	findBooks = append(findBooks, fmt.Sprintf("ID: %d\nISBN: %s\nАвтор: %s\nНазвание: %s\nКоличество в библиотеке (шт): %d", book.ID, book.ISBN, book.Author, book.Name, book.Count))
	//}
	//h.builder.NewMessage(msg, fmt.Sprintf("Книги найдены, вот информация о них:\n\n%v\n\nЧтобы арендовать какую-то из этих книг, нажмите на кнопку с её ID.", strings.Join(findBooks, "\n\n")), h.keyboard.FindBy(msg.CallbackQuery.From.ID, "name", name.Text, ids...))
	//return
}

func (h *BooksHandler) Next(ctx context.Context, msg telego.Update) {
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

	var findType string
	var searched string

	var arr []string

	jsonparser.ArrayEach([]byte(msg.CallbackQuery.Data), func(value []byte, dataType jsonparser.ValueType, offset int, err error) {
		if err != nil {
			fmt.Println(err)
			return
		}

		if dataType == jsonparser.String {
			arr = append(arr, string(value))
		}
	}, "searched")

	findType = arr[0]
	searched = arr[1]

	if err != nil {
		h.builder.NewMessage(msg, "Попробуйте заново.", h.keyboard.FindBook())
		slog.Error(err.Error())
		return
	}

	switch findType {
	case "author":
		h.Fetches[msg.CallbackQuery.From.ID] += offset + 9
		dto := dto2.NewAuthorInput(searched, h.Fetches[msg.CallbackQuery.From.ID])
		books, err := h.service.FindByAuthor(ctx, dto)
		if err != nil || books.Status != 200 {
			if err.Error() == "rpc error: code = Unknown desc = book is not found" {
				h.builder.NewMessage(msg, "Больше книг от этого автора не существует.", h.keyboard.FindBook())
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
			findBooks = append(findBooks, fmt.Sprintf("ID: %d\nISBN: %s\nАвтор: %s\nНазвание: %s\nКоличество в библиотеке (шт): %d", book.ID, book.ISBN, book.Author, book.Name, book.Count))
		}
		var predOffset int64
		if h.Fetches[msg.CallbackQuery.From.ID] >= 18 {
			predOffset = h.Fetches[msg.CallbackQuery.From.ID] - 9*2
		} else {
			predOffset = h.Fetches[msg.CallbackQuery.From.ID] - 9
		}
		h.builder.NewMessage(msg, fmt.Sprintf("Книги найдены, вот информация о них:\n\n%v\n\nЧтобы арендовать какую-то из этих книг, нажмите на кнопку с её ID.", strings.Join(findBooks, "\n\n")), h.keyboard.FindBy(findType, int(h.Fetches[msg.CallbackQuery.From.ID]), int(predOffset), searched, ids...))
		return
	case "name":
		h.Fetches[msg.CallbackQuery.From.ID] += 9
		dto := dto2.NewNameInput(searched, h.Fetches[msg.CallbackQuery.From.ID])
		books, err := h.service.FindByName(ctx, dto)
		if err != nil || books.Status != 200 {
			if err.Error() == "rpc error: code = Unknown desc = book is not found" {
				h.builder.NewMessage(msg, "Больше книг от этого автора не существует.", h.keyboard.FindBook())
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
			findBooks = append(findBooks, fmt.Sprintf("ID: %d\nISBN: %s\nАвтор: %s\nНазвание: %s\nКоличество в библиотеке (шт): %d", book.ID, book.ISBN, book.Author, book.Name, book.Count))
		}
		var predOffset int64
		if h.Fetches[msg.CallbackQuery.From.ID] >= 18 {
			predOffset = h.Fetches[msg.CallbackQuery.From.ID] - 9*2
		} else {
			predOffset = h.Fetches[msg.CallbackQuery.From.ID] - 9
		}
		h.builder.NewMessage(msg, fmt.Sprintf("Книги найдены, вот информация о них:\n\n%v\n\nЧтобы арендовать какую-то из этих книг, нажмите на кнопку с её ID.", strings.Join(findBooks, "\n\n")), h.keyboard.FindBy(findType, int(h.Fetches[msg.CallbackQuery.From.ID]), int(predOffset), searched, ids...))
		return
	}
}
