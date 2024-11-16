package main

import "fmt"

func main() {
	fmt.Println(DontGiveMeFive(1, 9))  // Ожидаемый результат: 8
	fmt.Println(DontGiveMeFive(4, 17)) // Ожидаемый результат: 12
}

func DontGiveMeFive(start int, end int) int {
	count := 0
	for i := start; i <= end; i++ {
		if !containsFive(i) {
			count++
		}
	}
	return count
}

func containsFive(n int) bool {
	for n > 0 {
		if n%10 == 5 {
			return true
		}
		n /= 10
	}
	return false
}
