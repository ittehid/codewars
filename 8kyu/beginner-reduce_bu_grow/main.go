package main

import "fmt"

func main() {
	arr1 := []int{1, 2, 3}
	arr2 := []int{4, 1, 1, 1, 4}
	arr3 := []int{2, 2, 2, 2, 2, 2}

	fmt.Println(Grow(arr1))
	fmt.Println(Grow(arr2))
	fmt.Println(Grow(arr3))
}

func Grow(arr []int) int {
	result := 1
	for _, val := range arr {
		result *= val
	}
	return result
}
