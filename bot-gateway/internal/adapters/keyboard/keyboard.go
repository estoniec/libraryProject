package keyboard

import (
	"encoding/json"
	"gateway/pkg/adapters/builder"
	"github.com/mymmrac/telego"
	tu "github.com/mymmrac/telego/telegoutil"
)

type KeyboardManager struct {
	Fetches map[int64]int
}

func NewKeyboardManager() *KeyboardManager {
	return &KeyboardManager{
		Fetches: make(map[int64]int),
	}
}

func (k *KeyboardManager) Repeat() *telego.InlineKeyboardMarkup {
	// TODO переделать роутер под это

	keyboard := tu.InlineKeyboard(
		tu.InlineKeyboardRow(
			tu.InlineKeyboardButton("Попробовать снова").WithCallbackData(addCommand("/start")),
		),
	)
	return keyboard
}

func (k *KeyboardManager) Menu() *telego.InlineKeyboardMarkup {
	keyboard := tu.InlineKeyboard(
		tu.InlineKeyboardRow(
			tu.InlineKeyboardButton("Взять книгу").WithCallbackData(addCommand("/available")),
		),
		tu.InlineKeyboardRow(
			tu.InlineKeyboardButton("Сколько я взял книг?").WithCallbackData(addCommand("/rented")),
		),
	)
	return keyboard
}

func (k *KeyboardManager) FindBook() *telego.InlineKeyboardMarkup {
	keyboard := tu.InlineKeyboard(
		tu.InlineKeyboardRow(
			tu.InlineKeyboardButton("Поиск по ISBN").WithCallbackData(addCommand("/findbyisbn")),
		),
		tu.InlineKeyboardRow(
			tu.InlineKeyboardButton("Поиск по автору").WithCallbackData(addCommand("/findbyauthor")),
			tu.InlineKeyboardButton("Поиск по названию").WithCallbackData(addCommand("/findbyname")),
		),
		tu.InlineKeyboardRow(
			tu.InlineKeyboardButton("Поиск по автору и названию").WithCallbackData(addCommand("/findbyauthorandname")),
		),
		tu.InlineKeyboardRow(
			tu.InlineKeyboardButton("Поиск всех книг").WithCallbackData(addCommand("/findall")),
		),
	)
	return keyboard
}

func (k *KeyboardManager) FindBy(fromID int64, findType string, something string, id ...string) *telego.InlineKeyboardMarkup {
	offsetNext := 0
	offsetPred := 0
	if _, ok := k.Fetches[fromID]; !ok {
		k.Fetches[fromID] = 0
	} else {
		k.Fetches[fromID] += 10
		offsetNext = k.Fetches[fromID]
		offsetPred = k.Fetches[fromID] - 10
	}
	jsonNext, _ := json.Marshal(builder.NewPayload().SetCommand("/nextAuthor").AddPayload("offset", offsetNext).AddPayload(findType, something))
	jsonPred, _ := json.Marshal(builder.NewPayload().SetCommand("/predAuthor").AddPayload("offset", offsetPred).AddPayload(findType, something))
	var row []telego.InlineKeyboardButton
	var rows [][]telego.InlineKeyboardButton
	for i := 0; i < len(id); i++ {
		button := tu.InlineKeyboardButton(id[i]).WithCallbackData(addPayloadForRentButtons(id[i]))
		row = append(row, button)
		if len(row) == 3 || i == len(id)-1 {
			rows = append(rows, row)
			row = make([]telego.InlineKeyboardButton, 0)
		}
	}
	keyboard := tu.InlineKeyboard(
		rows[0],
		rows[1],
		rows[2],
		tu.InlineKeyboardRow(
			tu.InlineKeyboardButton("Следующая страница").WithCallbackData(string(jsonNext)),
			tu.InlineKeyboardButton("Предыдущая страница").WithCallbackData(string(jsonPred)),
		),
	)
	return keyboard
}

func addPayloadForRentButtons(id string) string {
	json, _ := json.Marshal(builder.NewPayload().SetCommand("/rent").AddPayload("id", id))
	return string(json)
}

func addCommand(command string) string {
	json, _ := json.Marshal(builder.NewPayload().SetCommand(command))
	return string(json)
}
