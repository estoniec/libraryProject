package utils

import "regexp"

func IsPhoneNumber(phoneNumber string) bool {
	// Паттерн для проверки формата телефонного номера
	pattern := "^\\+[1-9]\\d{1,14}$"
	match, _ := regexp.MatchString(pattern, phoneNumber)
	return match
}

func IsNameAndSurname(username string) bool {
	re := regexp.MustCompile("^[А-ЯЁ][а-яё]+\\s[А-ЯЁ][а-яё]+$")
	return re.MatchString(username)
}

func IsValidClassAndParallel(str string) bool {
	re := regexp.MustCompile("^[1-9][0-1]?[А-Г]$")
	return re.MatchString(str)
}
