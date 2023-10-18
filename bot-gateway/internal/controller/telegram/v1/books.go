package v1

import (
	"context"
	"fmt"
	"gateway/internal/controller/telegram/dto"
	books_dto "gateway/internal/domain/books/dto"
	"gateway/internal/domain/books/model"
	"gateway/pkg/adapters/handling"
	"github.com/buger/jsonparser"
	"github.com/mymmrac/telego"
	"log/slog"
	"strconv"
	"strings"
	"time"
)

type BooksUsecase interface {
	FindBy(ctx context.Context, input dto.FindByInput) (books_dto.FindByOutput, error)
	CreateSearch(ctx context.Context, input dto.CreateSearchInput) (books_dto.CreateSearchOutput, error)
	FindSearch(ctx context.Context, input dto.FindSearchInput) (books_dto.FindSearchOutput, error)
}

type BooksKeyboard interface {
	FindBook() *telego.InlineKeyboardMarkup
	FindBy(offsetNext int, offsetPred int, getID int64, id ...string) *telego.InlineKeyboardMarkup
	Find(id string) *telego.InlineKeyboardMarkup
	Menu() *telego.InlineKeyboardMarkup
}

type BooksHandler struct {
	builder          Builder
	router           Router
	question         Question
	callbackQuestion CallbackQuestion
	usecase          BooksUsecase
	keyboard         BooksKeyboard
}

func NewBooksHandler(builder Builder, router Router, question Question, callbackQuestion CallbackQuestion, usecase BooksUsecase, keyboard BooksKeyboard) *BooksHandler {
	return &BooksHandler{
		builder:          builder,
		router:           router,
		question:         question,
		callbackQuestion: callbackQuestion,
		usecase:          usecase,
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
	findByName := regGroup.NewHandler(h.FindByName)
	findByName.WithCommand("/findbyname")
	findByNameAndAuthor := regGroup.NewHandler(h.FindByNameAndAuthor)
	findByNameAndAuthor.WithCommand("/findbynameandauthor")
	findAll := regGroup.NewHandler(h.FindAll)
	findAll.WithCommand("/findall")
	nextBtn := regGroup.NewHandler(h.Next)
	nextBtn.WithCommand("/next")
	predBtn := regGroup.NewHandler(h.Pred)
	predBtn.WithCommand("/pred")
	menu := regGroup.NewHandler(h.Menu)
	menu.WithCommand("/menu")
	h.AddGroup(regGroup)
}

func (h *BooksHandler) Menu(ctx context.Context, msg telego.Update) {
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
	return
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
	dto := dto.NewByInput(0, model.NewFindBook(isbn.Text, "", ""))
	book, err := h.usecase.FindBy(ctx, dto)
	if err != nil || book.Status != 200 {
		if err.Error() == "rpc error: code = Unknown desc = book is not found" {
			h.builder.NewMessage(msg, "Такой книги не в библиотеке существует.", h.keyboard.FindBook())
			return
		}
		h.builder.NewMessage(msg, "Попробуйте заново.", h.keyboard.FindBook())
		slog.Error(err.Error())
		return
	}
	var findBooks []string
	var ids []string
	for _, book := range book.Book {
		ids = append(ids, strconv.Itoa(book.ID))
		findBooks = append(findBooks, fmt.Sprintf("ID: %d\nISBN: %s\nАвтор: %s\nНазвание: %s\nКоличество в библиотеке (шт): %d", book.ID, book.ISBN, book.Author, book.Name, book.Count))
	}
	h.builder.NewMessage(msg, fmt.Sprintf("Книги найдены, вот информация о них:\n\n%v\n\nЧтобы арендовать какую-то из этих книг, нажмите на кнопку с её ID.", strings.Join(findBooks, "\n\n")), h.keyboard.Find(ids[0]))
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
	bookDTO := dto.NewByInput(0, model.NewFindBook("", "", author.Text))
	book, err := h.usecase.FindBy(ctx, bookDTO)
	if err != nil || book.Status != 200 {
		if err.Error() == "rpc error: code = Unknown desc = book is not found" {
			h.builder.NewMessage(msg, "Такой книги не в библиотеке существует.", h.keyboard.FindBook())
			return
		}
		h.builder.NewMessage(msg, "Попробуйте заново.", h.keyboard.FindBook())
		slog.Error(err.Error())
		return
	}
	var findBooks []string
	var ids []string
	for _, book := range book.Book {
		ids = append(ids, strconv.Itoa(book.ID))
		findBooks = append(findBooks, fmt.Sprintf("ID: %d\nISBN: %s\nАвтор: %s\nНазвание: %s\nКоличество в библиотеке (шт): %d", book.ID, book.ISBN, book.Author, book.Name, book.Count))
	}
	ID := time.Now().Unix()
	createDTO := dto.NewCreateInput(ID, "author", author.Text)
	res, err := h.usecase.CreateSearch(ctx, createDTO)
	if err != nil {
		h.builder.NewMessage(msg, "Попробуйте заново.", h.keyboard.FindBook())
		slog.Error(res.Error)
		return
	}
	h.builder.NewMessage(msg, fmt.Sprintf("Книги найдены, вот информация о них:\n\n%v\n\nЧтобы арендовать какую-то из этих книг, нажмите на кнопку с её ID.", strings.Join(findBooks, "\n\n")), h.keyboard.FindBy(9, 0, ID, ids...))
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
	dtoBy := dto.NewByInput(0, model.NewFindBook("", name.Text, ""))
	book, err := h.usecase.FindBy(ctx, dtoBy)
	if err != nil || book.Status != 200 {
		if err.Error() == "rpc error: code = Unknown desc = book is not found" {
			h.builder.NewMessage(msg, "Такой книги не в библиотеке существует.", h.keyboard.FindBook())
			return
		}
		h.builder.NewMessage(msg, "Попробуйте заново.", h.keyboard.FindBook())
		slog.Error(err.Error())
		return
	}
	var findBooks []string
	var ids []string
	for _, book := range book.Book {
		ids = append(ids, strconv.Itoa(book.ID))
		findBooks = append(findBooks, fmt.Sprintf("ID: %d\nISBN: %s\nАвтор: %s\nНазвание: %s\nКоличество в библиотеке (шт): %d", book.ID, book.ISBN, book.Author, book.Name, book.Count))
	}
	ID := time.Now().Unix()
	createDTO := dto.NewCreateInput(ID, "name", name.Text)
	res, err := h.usecase.CreateSearch(ctx, createDTO)
	if err != nil {
		h.builder.NewMessage(msg, "Попробуйте заново.", h.keyboard.FindBook())
		slog.Error(res.Error)
		return
	}
	h.builder.NewMessage(msg, fmt.Sprintf("Книги найдены, вот информация о них:\n\n%v\n\nЧтобы арендовать какую-то из этих книг, нажмите на кнопку с её ID.", strings.Join(findBooks, "\n\n")), h.keyboard.FindBy(9, 0, ID, ids...))
	return
}

func (h *BooksHandler) FindByNameAndAuthor(ctx context.Context, msg telego.Update) {
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
	_, err = h.builder.NewMessage(msg, "Введите название книги (например, \"Россия\"):", nil)
	if err != nil {
		h.builder.NewMessage(msg, "Попробуйте заново.", h.keyboard.FindBook())
		slog.Error(err.Error())
		return
	}
	name, ok := <-answer
	if !ok || name.Text == "" {
		h.builder.NewMessage(msg, "Попробуйте заново.", h.keyboard.FindBook())
		return
	}
	dto := dto.NewByInput(0, model.NewFindBook("", name.Text, author.Text))
	book, err := h.usecase.FindBy(ctx, dto)
	if err != nil || book.Status != 200 {
		if err.Error() == "rpc error: code = Unknown desc = book is not found" {
			h.builder.NewMessage(msg, "Такой книги не в библиотеке существует.", h.keyboard.FindBook())
			return
		}
		h.builder.NewMessage(msg, "Попробуйте заново.", h.keyboard.FindBook())
		slog.Error(err.Error())
		return
	}
	var findBooks []string
	var ids []string
	for _, book := range book.Book {
		ids = append(ids, strconv.Itoa(book.ID))
		findBooks = append(findBooks, fmt.Sprintf("ID: %d\nISBN: %s\nАвтор: %s\nНазвание: %s\nКоличество в библиотеке (шт): %d", book.ID, book.ISBN, book.Author, book.Name, book.Count))
	}
	h.builder.NewMessage(msg, fmt.Sprintf("Книги найдены, вот информация о них:\n\n%v\n\nЧтобы арендовать какую-то из этих книг, нажмите на кнопку с её ID.", strings.Join(findBooks, "\n\n")), h.keyboard.Find(ids[0]))
	return
}

func (h *BooksHandler) FindAll(ctx context.Context, msg telego.Update) {
	err := h.builder.NewCallbackMessage(msg.CallbackQuery, "")
	if err != nil {
		h.builder.NewMessage(msg, "Попробуйте заново.", h.keyboard.FindBook())
		slog.Error(err.Error())
		return
	}
	dtoBy := dto.NewByInput(0, model.NewFindBook("", "", ""))
	book, err := h.usecase.FindBy(ctx, dtoBy)
	if err != nil || book.Status != 200 {
		if err.Error() == "rpc error: code = Unknown desc = book is not found" {
			h.builder.NewMessage(msg, "Такой книги не в библиотеке существует.", h.keyboard.FindBook())
			return
		}
		h.builder.NewMessage(msg, "Попробуйте заново.", h.keyboard.FindBook())
		slog.Error(err.Error())
		return
	}
	var findBooks []string
	var ids []string
	for _, book := range book.Book {
		ids = append(ids, strconv.Itoa(book.ID))
		findBooks = append(findBooks, fmt.Sprintf("ID: %d\nISBN: %s\nАвтор: %s\nНазвание: %s\nКоличество в библиотеке (шт): %d", book.ID, book.ISBN, book.Author, book.Name, book.Count))
	}
	ID := time.Now().Unix()
	createDTO := dto.NewCreateInput(ID, "all", "")
	res, err := h.usecase.CreateSearch(ctx, createDTO)
	if err != nil {
		h.builder.NewMessage(msg, "Попробуйте заново.", h.keyboard.FindBook())
		slog.Error(res.Error)
		return
	}
	h.builder.NewMessage(msg, fmt.Sprintf("Книги найдены, вот информация о них:\n\n%v\n\nЧтобы арендовать какую-то из этих книг, нажмите на кнопку с её ID.", strings.Join(findBooks, "\n\n")), h.keyboard.FindBy(9, 0, ID, ids...))
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

	getID, err := jsonparser.GetInt([]byte(msg.CallbackQuery.Data), "get")
	if err != nil {
		h.builder.NewMessage(msg, "Попробуйте заново.", h.keyboard.FindBook())
		slog.Error(err.Error())
		return
	}

	findDTO := dto.NewFindInput(getID)
	find, err := h.usecase.FindSearch(ctx, findDTO)
	if err != nil {
		h.builder.NewMessage(msg, "Попробуйте заново.", h.keyboard.FindBook())
		slog.Error(find.Error)
		return
	}

	switch find.Searched[0] {
	case "author":
		dtoBy := dto.NewByInput(offset, model.NewFindBook("", "", find.Searched[1]))
		books, err := h.usecase.FindBy(ctx, dtoBy)
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
		ID := time.Now().Unix()
		createDTO := dto.NewCreateInput(ID, "author", find.Searched[1])
		res, err := h.usecase.CreateSearch(ctx, createDTO)
		if err != nil {
			h.builder.NewMessage(msg, "Попробуйте заново.", h.keyboard.FindBook())
			slog.Error(res.Error)
			return
		}
		var predOffset int64
		predOffset = offset - 9
		h.builder.NewMessage(msg, fmt.Sprintf("Книги найдены, вот информация о них:\n\n%v\n\nЧтобы арендовать какую-то из этих книг, нажмите на кнопку с её ID.", strings.Join(findBooks, "\n\n")), h.keyboard.FindBy(int(offset+9), int(predOffset), ID, ids...))
		return
	case "name":
		dtoBy := dto.NewByInput(offset, model.NewFindBook("", find.Searched[1], ""))
		books, err := h.usecase.FindBy(ctx, dtoBy)
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
		ID := time.Now().Unix()
		createDTO := dto.NewCreateInput(ID, "name", find.Searched[1])
		res, err := h.usecase.CreateSearch(ctx, createDTO)
		if err != nil {
			h.builder.NewMessage(msg, "Попробуйте заново.", h.keyboard.FindBook())
			slog.Error(res.Error)
			return
		}
		var predOffset int64
		predOffset = offset - 9
		h.builder.NewMessage(msg, fmt.Sprintf("Книги найдены, вот информация о них:\n\n%v\n\nЧтобы арендовать какую-то из этих книг, нажмите на кнопку с её ID.", strings.Join(findBooks, "\n\n")), h.keyboard.FindBy(int(offset+9), int(predOffset), ID, ids...))
		return
	case "all":
		dtoBy := dto.NewByInput(offset, model.NewFindBook("", "", ""))
		books, err := h.usecase.FindBy(ctx, dtoBy)
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
		ID := time.Now().Unix()
		createDTO := dto.NewCreateInput(ID, "all", find.Searched[1])
		res, err := h.usecase.CreateSearch(ctx, createDTO)
		if err != nil {
			h.builder.NewMessage(msg, "Попробуйте заново.", h.keyboard.FindBook())
			slog.Error(res.Error)
			return
		}
		var predOffset int64
		predOffset = offset - 9
		h.builder.NewMessage(msg, fmt.Sprintf("Книги найдены, вот информация о них:\n\n%v\n\nЧтобы арендовать какую-то из этих книг, нажмите на кнопку с её ID.", strings.Join(findBooks, "\n\n")), h.keyboard.FindBy(int(offset+9), int(predOffset), ID, ids...))
		return
	}
}

func (h *BooksHandler) Pred(ctx context.Context, msg telego.Update) {
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

	getID, err := jsonparser.GetInt([]byte(msg.CallbackQuery.Data), "get")
	if err != nil {
		h.builder.NewMessage(msg, "Попробуйте заново.", h.keyboard.FindBook())
		slog.Error(err.Error())
		return
	}

	findDTO := dto.NewFindInput(getID)
	find, err := h.usecase.FindSearch(ctx, findDTO)
	if err != nil {
		h.builder.NewMessage(msg, "Попробуйте заново.", h.keyboard.FindBook())
		slog.Error(find.Error)
		return
	}

	switch find.Searched[0] {
	case "author":
		dtoBy := dto.NewByInput(offset, model.NewFindBook("", "", find.Searched[1]))
		books, err := h.usecase.FindBy(ctx, dtoBy)
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
		ID := time.Now().Unix()
		createDTO := dto.NewCreateInput(ID, "author", find.Searched[1])
		res, err := h.usecase.CreateSearch(ctx, createDTO)
		if err != nil {
			h.builder.NewMessage(msg, "Попробуйте заново.", h.keyboard.FindBook())
			slog.Error(res.Error)
			return
		}
		var predOffset int64
		if offset > 0 {
			predOffset = offset - 9
		} else {
			predOffset = 0
		}
		h.builder.NewMessage(msg, fmt.Sprintf("Книги найдены, вот информация о них:\n\n%v\n\nЧтобы арендовать какую-то из этих книг, нажмите на кнопку с её ID.", strings.Join(findBooks, "\n\n")), h.keyboard.FindBy(int(offset+9), int(predOffset), ID, ids...))
		return
	case "name":
		dtoBy := dto.NewByInput(offset, model.NewFindBook("", find.Searched[1], ""))
		books, err := h.usecase.FindBy(ctx, dtoBy)
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
		ID := time.Now().Unix()
		createDTO := dto.NewCreateInput(ID, "name", find.Searched[1])
		res, err := h.usecase.CreateSearch(ctx, createDTO)
		if err != nil {
			h.builder.NewMessage(msg, "Попробуйте заново.", h.keyboard.FindBook())
			slog.Error(res.Error)
			return
		}
		var predOffset int64
		if offset > 0 {
			predOffset = offset - 9
		} else {
			predOffset = 0
		}
		h.builder.NewMessage(msg, fmt.Sprintf("Книги найдены, вот информация о них:\n\n%v\n\nЧтобы арендовать какую-то из этих книг, нажмите на кнопку с её ID.", strings.Join(findBooks, "\n\n")), h.keyboard.FindBy(int(offset+9), int(predOffset), ID, ids...))
		return
	case "all":
		dtoBy := dto.NewByInput(offset, model.NewFindBook("", "", ""))
		books, err := h.usecase.FindBy(ctx, dtoBy)
		if err != nil || books.Status != 200 {
			if err.Error() == "rpc error: code = Unknown desc = book is not found" {
				h.builder.NewMessage(msg, "Больше книг в библиотеке нет.", h.keyboard.FindBook())
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
		ID := time.Now().Unix()
		createDTO := dto.NewCreateInput(ID, "author", find.Searched[1])
		res, err := h.usecase.CreateSearch(ctx, createDTO)
		if err != nil {
			h.builder.NewMessage(msg, "Попробуйте заново.", h.keyboard.FindBook())
			slog.Error(res.Error)
			return
		}
		var predOffset int64
		if offset > 0 {
			predOffset = offset - 9
		} else {
			predOffset = 0
		}
		h.builder.NewMessage(msg, fmt.Sprintf("Книги найдены, вот информация о них:\n\n%v\n\nЧтобы арендовать какую-то из этих книг, нажмите на кнопку с её ID.", strings.Join(findBooks, "\n\n")), h.keyboard.FindBy(int(offset+9), int(predOffset), ID, ids...))
		return
	}
}
