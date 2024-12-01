package main

import "fmt"

func main() {
	arr1 := []int{1, 2, 10, 50, 5}
	fmt.Println(SortNumbers(arr1))
}

func SortNumbers(numbers []int) []int {
	if numbers == nil || len(numbers) == 0 {
		return numbers
	} else {
		n := len(numbers)
		for i := 0; i < n; i++ {
			for j := 0; j < n-i-1; j++ {
				if numbers[j] > numbers[j+1] {
					numbers[j], numbers[j+1] = numbers[j+1], numbers[j]
				}
			}
		}
	}
	return numbers
}
