package main

import (
	"fmt"
	"strings"
)

func main() {
	fmt.Println(ReverseWords("This is an example!"))
	fmt.Println(ReverseWords("double  spaces"))
	fmt.Println(ReverseWords(".god yzal eht revo spmuj xof nworb kciuq eh"))
}

func ReverseWords(str string) string {
	result := []rune{}                // Результирующий срез рун (символов)
	newStr := strings.Split(str, " ") // Разбиваем исходную строку на слова по пробелу
	for i := 0; i < len(newStr); i++ {
		temp := []rune(newStr[i])             // Преобразуем текущее слово в срез рун
		for j := len(temp) - 1; j >= 0; j-- { // Проходим по символам слова в обратном порядке
			result = append(result, temp[j]) // Добавляем текущий символ в результат
		}
		if i < len(newStr)-1 { // Если это не последнее слово
			result = append(result, ' ') // Добавляем пробел после слова
		}
	}
	return string(result) // Преобразуем срез рун обратно в строку и возвращаем результат
}
