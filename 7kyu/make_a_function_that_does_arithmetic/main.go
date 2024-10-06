package main

import "fmt"

func main() {
	fmt.Println(Arithmetic(5, 2, "add"))      //+
	fmt.Println(Arithmetic(5, 2, "subtract")) //-
	fmt.Println(Arithmetic(5, 2, "multiply")) //*
	fmt.Println(Arithmetic(5, 2, "divide"))   //\
}

func Arithmetic(a int, b int, operator string) int {
	switch operator {
	case "add":
		return a + b
	case "subtract":
		return a - b
	case "multiply":
		return a * b
	case "divide":
		return a / b
	}
	return 0
}
