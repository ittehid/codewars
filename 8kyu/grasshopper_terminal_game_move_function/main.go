package main

import "fmt"

func main() {
	fmt.Println(Move(3, 6))
	fmt.Println(Move(0, 4))
	fmt.Println(Move(2, 5))
}

func Move(position int, roll int) int {
	return position + roll*2
}
