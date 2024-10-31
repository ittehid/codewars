package main

import (
	"fmt"
	"strings"
	"unicode"
)

func main() {
	fmt.Println(solve("a"))
	fmt.Println(solve("Z"))
	fmt.Println(solve("coDe"))
	fmt.Println(solve("CODe"))
	fmt.Println(solve("coD"))
	fmt.Println(solve("abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ"))
}

func solve(str string) string {
	titleCount := 0
	lowerCount := 0
	for _, val := range str {
		if unicode.IsUpper(val) {
			titleCount++
		} else {
			lowerCount++
		}
	}

	if titleCount == lowerCount {
		return strings.ToLower(str)
	} else if titleCount > lowerCount {
		return strings.ToUpper(str)
	}

	return strings.ToLower(str)
}
