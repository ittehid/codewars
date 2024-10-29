package main

import "fmt"

func main() {
	fmt.Println(Contamination("abc", "z"))
	fmt.Println(Contamination("", "z"))
	fmt.Println(Contamination("abc", ""))
	fmt.Println(Contamination("_3ebzgh4", "&"))
	fmt.Println(Contamination("//case", " "))
}

func Contamination(text, char string) string {
	if char == "" || text == "" {
		return ""
	} else if char == " " {
		for range text {
			return RangeText(text, " ")
		}
	} else {
		for range text {
			return RangeText(text, char)
		}
	}
	return ""
}

func RangeText(text, char string) string {
	newStr := ""
	for range text {
		newStr += char
	}
	return newStr
}
