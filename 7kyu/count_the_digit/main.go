package main

import (
	"fmt"
	"strconv"
)

func main() {
	fmt.Println(NbDig(550, 5))
	fmt.Println(NbDig(5750, 0))
}

func NbDig(n int, d int) int {
	count := 0

	for i := 0; i <= n; i++ {
		numSquare := strconv.Itoa(i * i)
		for _, num := range numSquare {
			if string(num) == strconv.Itoa(d) {
				count++
			}
		}
	}
	return count
}
