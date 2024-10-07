package main

import (
	"fmt"
	"strconv"
)

func main() {
	fmt.Println(StringToNumber("1234"))
	fmt.Println(StringToNumber("605"))
	fmt.Println(StringToNumber("1405"))
}

func StringToNumber(str string) int {
	result, err := strconv.Atoi(str)
	if err != nil {
		fmt.Println(err)
	}
	return result
}
