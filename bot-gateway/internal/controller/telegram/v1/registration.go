package v1

import (
	"context"
	dto2 "gateway/internal/domain/users/dto"
	"gateway/pkg/adapters/handling"
	"gateway/pkg/utils"
	"github.com/mymmrac/telego"
	"log/slog"
)

type RegService interface {
	Registration(context context.Context, input dto2.RegInput) (dto2.RegOutput, error)
	CheckUser(ctx context.Context, input dto2.CheckInput) (dto2.CheckOutput, error)
}

type RegKeyboard interface {
	Repeat() *telego.InlineKeyboardMarkup
	Menu() *telego.InlineKeyboardMarkup
}

type RegHandler struct {
	builder          Builder
	router           Router
	question         Question
	callbackQuestion CallbackQuestion
	service          RegService
	keyboard         RegKeyboard
}

func NewRegHandler(builder Builder, router Router, question Question, callbackQuestion CallbackQuestion, service RegService, keyboard RegKeyboard) *RegHandler {
	return &RegHandler{
		builder:          builder,
		router:           router,
		question:         question,
		callbackQuestion: callbackQuestion,
		service:          service,
		keyboard:         keyboard,
	}
}

func (h *RegHandler) AddGroup(handlerGroup handling.HandlersGroup) {
	h.router.AddGroup(handlerGroup)
}

func (h *RegHandler) Register() {
	regGroup := handling.NewHandlersGroup()
	reg := regGroup.NewHandler(h.Registration)
	reg.WithCommand("/start")
	h.AddGroup(regGroup)
}

func (h *RegHandler) Registration(ctx context.Context, msg telego.Update) {
	var checkDTO dto2.CheckInput

	if msg.CallbackQuery != nil {
		h.builder.NewCallbackMessage(msg.CallbackQuery, "")
		checkDTO = dto2.NewCheckInput(msg.CallbackQuery.From.ID)
	} else {
		checkDTO = dto2.NewCheckInput(msg.Message.From.ID)
	}

	if checkedUser, err := h.service.CheckUser(ctx, checkDTO); checkedUser.Checked {
		if err != nil {
			slog.Error(err.Error())
			return
		}
		_, err = h.builder.NewMessage(msg, "Вы уже зарегистрированы в системе, вот ваш функционал:", h.keyboard.Menu())
		if err != nil {
			slog.Error(err.Error())
			return
		}
		return
	}

	_, err := h.builder.NewMessage(msg, "Здравствуйте!\nДля продолжения работы с ботом необходимо зарегистрироваться.\nДля начала, пожалуйста, введите свой контактный номер телефона:", nil)
	if err != nil {
		slog.Error(err.Error())
		return
	}
	answers, c := h.question.NewQuestion(msg)
	defer c()
	telephone, ok := <-answers
	if !ok || !utils.IsPhoneNumber(telephone.Text) {
		h.builder.NewMessage(msg, "Попробуйте ввести номер телефона заново.", h.keyboard.Repeat())
		return
	}

	_, err = h.builder.NewMessage(msg, "Введите своё имя и фамилию (через пробел, с большой буквы)", nil)
	if err != nil {
		slog.Error(err.Error())
		return
	}
	username, ok := <-answers
	if !ok || !utils.IsNameAndSurname(username.Text) {
		h.builder.NewMessage(msg, "Попробуйте ввести имя и фамилию заново.", h.keyboard.Repeat())
		return
	}

	_, err = h.builder.NewMessage(msg, "Введите свой класс и параллель (например, \"10A\")", nil)
	if err != nil {
		slog.Error(err.Error())
		return
	}
	class, ok := <-answers
	if !ok || !utils.IsValidClassAndParallel(class.Text) {
		h.builder.NewMessage(msg, "Попробуйте ввести класс и параллель заново.", h.keyboard.Repeat())
		return
	}
	dto := dto2.NewRegInput(telephone.Text, username.Text, class.Text, msg.Message.From.ID)
	res, err := h.service.Registration(ctx, dto)
	if err != nil || res.Error != "" {
		slog.Error(err.Error(), res.Error)
		h.builder.NewMessage(msg, "Возникла ошибка, повторите попытку позже.", h.keyboard.Repeat())
		return
	}
	h.builder.NewMessage(msg, "Вы успешно зарегистрированы в системе. Вот ваш функционал:", h.keyboard.Menu())
	return
}
