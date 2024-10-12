package main

import (
	"fmt"
	"strings"
)

func main() {
	fmt.Println(DuplicateEncode("din"))
	fmt.Println(DuplicateEncode("recede"))
	fmt.Println(DuplicateEncode("(( @"))
}

func DuplicateEncode(word string) string {
	arrMap := make(map[rune]int)
	word = strings.ToLower(word)
	result := ""

	for _, val := range word {
		arrMap[val]++
	}

	for _, val := range word {
		if arrMap[val] > 1 {
			result += ")"
		} else {
			result += "("
		}
	}
	return result
}
