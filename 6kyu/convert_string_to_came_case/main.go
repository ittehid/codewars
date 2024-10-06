package main

import (
	"fmt"
	"strings"
)

func main() {
	fmt.Println(ToCamelCase("the-stealth-warrior")) // theStealthWarrior
	fmt.Println(ToCamelCase("The_Stealth_Warrior")) // theStealthWarrior
	fmt.Println(ToCamelCase("The_Stealth-Warrior")) // theStealthWarrior
}

func ToCamelCase(s string) string {
	// Заменяем '-' и '_' на пробел, чтобы легко разбить строку на слова
	s = strings.ReplaceAll(s, "-", " ")
	s = strings.ReplaceAll(s, "_", " ")

	// Разбиваем строку на слова
	words := strings.Fields(s)

	// Создаем новый слайс для хранения слов в формате camelCase
	newWords := make([]string, len(words))

	// Перебираем слова и изменяем регистр первой буквы
	for i, word := range words {
		// Если это не первое слово, сделать первую букву заглавной
		if i > 0 {
			newWords[i] = strings.Title(word) // делает первую букву заглавной
		} else {
			newWords[i] = word // первое слово - с маленькой буквы
		}
	}

	// Объединяем слова в одну строку
	return strings.Join(newWords, "")
}
