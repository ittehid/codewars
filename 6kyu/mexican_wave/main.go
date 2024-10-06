package main

import (
	"fmt"
	"unicode"
)

func main() {
	fmt.Println(wave("hello"))
}

func wave(word string) []string {
	var result []string
	runes := []rune(word) // Работаем с рунами для поддержки Unicode

	for i, r := range runes {
		if unicode.IsLetter(r) {
			// Создаем копию исходного среза рун
			temp := make([]rune, len(runes))
			copy(temp, runes)

			// Делаем указанную букву заглавной
			temp[i] = unicode.ToUpper(r)

			// Добавляем новую строку в результат
			result = append(result, string(temp))
		}
	}
	return result
}
