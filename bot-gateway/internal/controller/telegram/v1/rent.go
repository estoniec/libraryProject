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
	bookID, err := jsonparser.GetString([]byte(msg.CallbackQuery.Data), "id")
	if err != nil {
		h.builder.NewMessage(msg, "Попробуйте заново.", h.keyboard.FindBook())
		slog.Error(err.Error())
		return
	}
	input := dto.NewRentInput(bookID, msg.CallbackQuery.From.ID, time.Now().Unix()+int64(daysInt*24*60*60))
	rent, err := h.usecase.RentBook(ctx, input)
	if err != nil || rent.Status == 404 {
		h.builder.NewMessage(msg, "Попробуйте заново.", h.keyboard.FindBook())
		slog.Error(err.Error())
		return
	}
	bookIDint, err := strconv.Atoi(bookID)
	if err != nil {
		h.builder.NewMessage(msg, fmt.Sprintf("Запрос на аренду книги создан, но информации о книге не найдено.\n\nЧтобы подтвердить аренду книги подойдите в библиотеку и продиктуйте номер: %d", rent.ID), h.keyboard.FindBook())
		return
	}
	inputBook := dto.NewByInput(0, model.NewFindBook("", "", "", int64(bookIDint)))
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
