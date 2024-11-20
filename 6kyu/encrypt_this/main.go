package main

import (
	"fmt"
	"strconv"
	"strings"
)

//Ваше сообщение представляет собой строку, содержащую слова, разделенные пробелами.
//Вам необходимо зашифровать каждое слово в сообщении, используя следующие правила:
//Первую букву необходимо преобразовать в ее код ASCII.
//Вторую букву необходимо поменять местами с последней буквой.

func main() {
	//fmt.Println(EncryptThis(""))
	fmt.Println(EncryptThis("A wise old owl lived in an oak"))
	fmt.Println(EncryptThis("The more he saw the less he spoke"))
}

func EncryptThis(text string) string {
	words := strings.Fields(text)
	var result []string
	for _, word := range words {
		if len(word) == 0 {
			continue
		}
		runeStr := []rune(word)
		asciiCode := strconv.Itoa(int(runeStr[0]))
		encryptedWord := asciiCode
		if len(runeStr) == 1 {
		} else if len(runeStr) == 2 {
			encryptedWord += string(runeStr[1])
		} else {
			runeStr[1], runeStr[len(runeStr)-1] = runeStr[len(runeStr)-1], runeStr[1]
			encryptedWord += string(runeStr[1:])
		}
		result = append(result, encryptedWord)
	}
	return strings.Join(result, " ")
}
