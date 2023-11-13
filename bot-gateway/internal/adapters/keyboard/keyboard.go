package keyboard

import (
	"encoding/json"
	"gateway/pkg/adapters/builder"
	"github.com/mymmrac/telego"
	tu "github.com/mymmrac/telego/telegoutil"
)

type KeyboardManager struct {
}

func NewKeyboardManager() *KeyboardManager {
	return &KeyboardManager{}
}

func (k *KeyboardManager) Repeat() *telego.InlineKeyboardMarkup {
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
			tu.InlineKeyboardButton("Поиск по автору и названию").WithCallbackData(addCommand("/findbynameandauthor")),
		),
		tu.InlineKeyboardRow(
			tu.InlineKeyboardButton("Поиск всех книг").WithCallbackData(addCommand("/findall")),
		),
		tu.InlineKeyboardRow(
			tu.InlineKeyboardButton("Отменить").WithCallbackData(addCommand("/menu")),
		),
	)
	return keyboard
}

func (k *KeyboardManager) FindBy(offsetNext int, offsetPred int, getID int64, id ...string) *telego.InlineKeyboardMarkup {
	jsonNext, _ := json.Marshal(builder.NewPayload().SetCommand("/next").AddPayload("get", getID).AddPayload("offset", offsetNext))
	jsonPred, _ := json.Marshal(builder.NewPayload().SetCommand("/pred").AddPayload("get", getID).AddPayload("offset", offsetPred))
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
	keyboard := tu.InlineKeyboard()

	if len(rows) > 0 {
		keyboard.InlineKeyboard = append(keyboard.InlineKeyboard, rows[0])
	}

	if len(rows) > 1 {
		keyboard.InlineKeyboard = append(keyboard.InlineKeyboard, rows[1])
	}

	if len(rows) > 2 {
		keyboard.InlineKeyboard = append(keyboard.InlineKeyboard, rows[2])
	}

	keyboard.InlineKeyboard = append(keyboard.InlineKeyboard,
		tu.InlineKeyboardRow(
			tu.InlineKeyboardButton("<<").WithCallbackData(string(jsonPred)),
			tu.InlineKeyboardButton(">>").WithCallbackData(string(jsonNext)),
		),
		tu.InlineKeyboardRow(
			tu.InlineKeyboardButton("Отменить").WithCallbackData(addCommand("/menu")),
		),
	)
	return keyboard
}

func (k *KeyboardManager) Find(id string) *telego.InlineKeyboardMarkup {
	keyboard := tu.InlineKeyboard(
		tu.InlineKeyboardRow(
			tu.InlineKeyboardButton(id).WithCallbackData(addPayloadForRentButtons(id)),
		),
		tu.InlineKeyboardRow(
			tu.InlineKeyboardButton("Отменить").WithCallbackData(addCommand("/menu")),
		),
	)
	return keyboard
}

func (k *KeyboardManager) Admin() *telego.ReplyKeyboardMarkup {
	keyboard := tu.Keyboard(
		tu.KeyboardRow(
			tu.KeyboardButton("Добавить книгу"),
		),
		tu.KeyboardRow(
			tu.KeyboardButton("Изменить количество книг"),
		),
		tu.KeyboardRow(
			tu.KeyboardButton("Удалить книгу"),
		),
		tu.KeyboardRow(
			tu.KeyboardButton("Подтвердить аренду книги"),
		),
		tu.KeyboardRow(
			tu.KeyboardButton("Подтвердить возврат книги"),
		),
	).WithResizeKeyboard()
	return keyboard
}

func (k *KeyboardManager) Success() *telego.InlineKeyboardMarkup {
	keyboard := tu.InlineKeyboard(
		tu.InlineKeyboardRow(
			tu.InlineKeyboardButton("Подтвердить").WithCallbackData(addCommand("/accept")),
		),
		tu.InlineKeyboardRow(
			tu.InlineKeyboardButton("Отменить").WithCallbackData(addCommand("/cancel")),
		),
	)
	return keyboard
}

func (k *KeyboardManager) PhoneNumber() *telego.ReplyKeyboardMarkup {
	keyboard := tu.Keyboard(
		tu.KeyboardRow(
			tu.KeyboardButton("Отправить номер телефона").WithRequestContact(),
		),
	).WithResizeKeyboard().WithOneTimeKeyboard()
	return keyboard
}

func (k *KeyboardManager) Switch(offsetNext int, offsetPred int, id ...string) *telego.InlineKeyboardMarkup {
	jsonNext, _ := json.Marshal(builder.NewPayload().SetCommand("/nextrents").AddPayload("offset", offsetNext))
	jsonPred, _ := json.Marshal(builder.NewPayload().SetCommand("/predrents").AddPayload("offset", offsetPred))
	var row []telego.InlineKeyboardButton
	var rows [][]telego.InlineKeyboardButton
	for i := 0; i < len(id); i++ {
		button := tu.InlineKeyboardButton(id[i]).WithCallbackData(addPayloadForReturnButtons(id[i]))
		row = append(row, button)
		if len(row) == 3 || i == len(id)-1 {
			rows = append(rows, row)
			row = make([]telego.InlineKeyboardButton, 0)
		}
	}
	keyboard := tu.InlineKeyboard()

	if len(rows) > 0 {
		keyboard.InlineKeyboard = append(keyboard.InlineKeyboard, rows[0])
	}

	if len(rows) > 1 {
		keyboard.InlineKeyboard = append(keyboard.InlineKeyboard, rows[1])
	}

	if len(rows) > 2 {
		keyboard.InlineKeyboard = append(keyboard.InlineKeyboard, rows[2])
	}

	keyboard.InlineKeyboard = append(keyboard.InlineKeyboard,
		tu.InlineKeyboardRow(
			tu.InlineKeyboardButton("<<").WithCallbackData(string(jsonPred)),
			tu.InlineKeyboardButton(">>").WithCallbackData(string(jsonNext)),
		),
		tu.InlineKeyboardRow(
			tu.InlineKeyboardButton("Отменить").WithCallbackData(addCommand("/menu")),
		),
	)
	return keyboard
}

func addPayloadForRentButtons(id string) string {
	json, _ := json.Marshal(builder.NewPayload().SetCommand("/rent").AddPayload("id", id))
	return string(json)
}

func addPayloadForReturnButtons(id string) string {
	json, _ := json.Marshal(builder.NewPayload().SetCommand("/returnbook").AddPayload("id", id))
	return string(json)
}

func addCommand(command string) string {
	json, _ := json.Marshal(builder.NewPayload().SetCommand(command))
	return string(json)
}
