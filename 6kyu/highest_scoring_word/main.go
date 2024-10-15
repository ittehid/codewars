package main

import (
	"fmt"
	"strings"
)

var (
	alphabetMap = map[string]int{
		"a": 1,
		"b": 2,
		"c": 3,
		"d": 4,
		"e": 5,
		"f": 6,
		"g": 7,
		"h": 8,
		"i": 9,
		"j": 10,
		"k": 11,
		"l": 12,
		"m": 13,
		"n": 14,
		"o": 15,
		"p": 16,
		"q": 17,
		"r": 18,
		"s": 19,
		"t": 20,
		"u": 21,
		"v": 22,
		"w": 23,
		"x": 24,
		"y": 25,
		"z": 26,
	}
)

func main() {
	fmt.Println(High("man i need a taxi up to ubud"))
	fmt.Println(High("what time are we climbing up the volcano"))
	fmt.Println(High("take me to semynak"))
	fmt.Println(High("aa b"))
	fmt.Println(High("b aa"))
	fmt.Println(High("bb d"))
	fmt.Println(High("d bb"))
	fmt.Println(High("aaa b"))
	fmt.Println(High("aaa b"))
}

func High(s string) string {
	result := ""
	tempMax := 0
	count := 0

	// переводим строку в слайс по пробелам
	slice := strings.Split(s, " ")
	// создаем карту для хранения результата
	resultMap := make(map[string]int)

	for _, val := range slice {
		count = 0
		for _, char := range val {
			count += alphabetMap[string(char)]
		}
		resultMap[val] += count
	}

	for key, val := range resultMap {
		if val > tempMax {
			tempMax = val
			result = key
		}
	}

	return result
}
