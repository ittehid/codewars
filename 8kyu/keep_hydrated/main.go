package main

import (
	"fmt"
)

func main() {
	fmt.Println(Litres(2))
	fmt.Println(Litres(0.82))
	fmt.Println(Litres(11.8))
}

func Litres(time float64) int {
	result := time * 0.5
	return int(result)
}
