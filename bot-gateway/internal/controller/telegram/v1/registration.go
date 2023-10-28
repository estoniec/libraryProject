package v1

import (
	"context"
	"gateway/internal/controller/telegram/dto"
	dto2 "gateway/internal/domain/users/dto"
	"gateway/pkg/adapters/handling"
	"gateway/pkg/utils"
	"github.com/mymmrac/telego"
	"log/slog"
)

type RegUsecase interface {
	Registration(context context.Context, input dto.RegInput) (dto2.RegOutput, error)
	CheckUser(ctx context.Context, input dto.CheckInput) (dto2.CheckOutput, error)
	CheckRole(ctx context.Context, input dto.CheckRoleInput) (dto2.CheckRoleOutput, error)
}

type RegKeyboard interface {
	Repeat() *telego.InlineKeyboardMarkup
	Menu() *telego.InlineKeyboardMarkup
	PhoneNumber() *telego.ReplyKeyboardMarkup
}

type RegHandler struct {
	builder          Builder
	router           Router
	question         Question
	callbackQuestion CallbackQuestion
	usecase          RegUsecase
	keyboard         RegKeyboard
}

func NewRegHandler(builder Builder, router Router, question Question, callbackQuestion CallbackQuestion, usecase RegUsecase, keyboard RegKeyboard) *RegHandler {
	return &RegHandler{
		builder:          builder,
		router:           router,
		question:         question,
		callbackQuestion: callbackQuestion,
		usecase:          usecase,
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
	var checkDTO dto.CheckInput

	if msg.CallbackQuery != nil {
		h.builder.NewCallbackMessage(msg.CallbackQuery, "")
		checkDTO = dto.NewCheckInput(msg.CallbackQuery.From.ID)
	} else {
		checkDTO = dto.NewCheckInput(msg.Message.From.ID)
	}

	if checkedUser, err := h.usecase.CheckUser(ctx, checkDTO); checkedUser.Checked {
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

	_, err := h.builder.NewMessageWithKeyboard(msg, "Здравствуйте!\nДля продолжения работы с ботом необходимо зарегистрироваться.\nДля начала, пожалуйста, введите свой контактный номер телефона с помощью кнопки (внизу)", h.keyboard.PhoneNumber())
	if err != nil {
		slog.Error(err.Error())
		return
	}
	answers, c := h.question.NewQuestion(msg)
	defer c()
	telephone, ok := <-answers
	if !ok {
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
	dto := dto.NewRegInput(telephone.Contact.PhoneNumber, username.Text, class.Text, msg.Message.From.ID)
	res, err := h.usecase.Registration(ctx, dto)
	if err != nil || res.Error != "" {
		slog.Error(err.Error(), res.Error)
		h.builder.NewMessage(msg, "Возникла ошибка, повторите попытку позже.", h.keyboard.Repeat())
		return
	}
	h.builder.NewMessage(msg, "Вы успешно зарегистрированы в системе. Вот ваш функционал:", h.keyboard.Menu())
	return
}
