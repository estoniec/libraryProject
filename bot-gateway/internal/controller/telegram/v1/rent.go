package v1

import (
	"context"
	"gateway/internal/controller/telegram/dto"
	"gateway/pkg/adapters/handling"
	"github.com/buger/jsonparser"
	"github.com/mymmrac/telego"
	"log/slog"
	"strconv"
	"time"
)

type RentUsecase interface{}

type RentKeyboard interface {
	FindBook() *telego.InlineKeyboardMarkup
}

type RentHandler struct {
	builder          Builder
	router           Router
	question         Question
	callbackQuestion CallbackQuestion
	rentUsecase      RentUsecase
	keyboard         RentKeyboard
}

func NewRentHandler(builder Builder, router Router, question Question, callback CallbackQuestion, rentUsecase RentUsecase, keyboard RentKeyboard) *RentHandler {
	return &RentHandler{
		builder:          builder,
		router:           router,
		question:         question,
		callbackQuestion: callback,
		rentUsecase:      rentUsecase,
		keyboard:         keyboard,
	}
}

func (h *RentHandler) AddGroup(handlerGroup handling.HandlersGroup) {
	h.router.AddGroup(handlerGroup)
}

func (h *RentHandler) Register() {
	regGroup := handling.NewHandlersGroup()

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
		h.builder.NewMessage(msg, "Попробуйте ввести количество книг заново.", nil)
		return
	}
	bookID, err := jsonparser.GetString([]byte(msg.CallbackQuery.Data), "id")
	if err != nil {
		h.builder.NewMessage(msg, "Попробуйте заново.", h.keyboard.FindBook())
		slog.Error(err.Error())
		return
	}
	input := dto.NewRentInput(bookID, msg.CallbackQuery.From.ID, time.Now().Unix()+int64(daysInt*24*60*60))

}
