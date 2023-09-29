package keyboard

import (
	"github.com/mymmrac/telego"
	tu "github.com/mymmrac/telego/telegoutil"
)

type KeyboardManager struct{}

func NewKeyboardManager() *KeyboardManager {
	return &KeyboardManager{}
}

func (k *KeyboardManager) Repeat() *telego.InlineKeyboardMarkup {
	keyboard := tu.InlineKeyboard(
		tu.InlineKeyboardRow(
			tu.InlineKeyboardButton("Попробовать снова").WithCallbackData("/start"),
		),
	)
	return keyboard
}

func (k *KeyboardManager) Menu() *telego.InlineKeyboardMarkup {
	keyboard := tu.InlineKeyboard(
		tu.InlineKeyboardRow(
			tu.InlineKeyboardButton("Взять книгу").WithCallbackData("/available"),
		),
		tu.InlineKeyboardRow(
			tu.InlineKeyboardButton("Сколько я взял книг?").WithCallbackData("/rented"),
		),
	)
	return keyboard
}

func (k *KeyboardManager) FindBook() *telego.InlineKeyboardMarkup {
	keyboard := tu.InlineKeyboard(
		tu.InlineKeyboardRow(
			tu.InlineKeyboardButton("Поиск по ISBN").WithCallbackData("/findbyISBN"),
		),
		tu.InlineKeyboardRow(
			tu.InlineKeyboardButton("Поиск по автору").WithCallbackData("/findbyauthor"),
			tu.InlineKeyboardButton("Поиск по названию").WithCallbackData("/findbyname"),
		),
		tu.InlineKeyboardRow(
			tu.InlineKeyboardButton("Поиск по автору и названию").WithCallbackData("/findbyauthorandname"),
		),
		tu.InlineKeyboardRow(
			tu.InlineKeyboardButton("Поиск всех книг").WithCallbackData("/findall"),
		),
	)
	return keyboard
}
