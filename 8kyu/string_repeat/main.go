package main

import "fmt"

func main() {
	fmt.Println(RepeatStr(4, "a"))
	fmt.Println(RepeatStr(3, "hello "))
	fmt.Println(RepeatStr(2, "abc"))
}

func RepeatStr(repetitions int, value string) string {
	str := ""
	for i := 0; i < repetitions; i++ {
		str += value
	}
	return str
}
