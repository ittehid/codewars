package main

import "fmt"

func main() {
	arr1 := [3]int{2, 3, 1}
	arr2 := [3]int{-15, -10, 14}
	arr3 := [3]int{15, 10, 14}
	arr4 := [3]int{2, 3, 1}

	fmt.Println(Gimme(arr1))
	fmt.Println(Gimme(arr2))
	fmt.Println(Gimme(arr3))
	fmt.Println(Gimme(arr4))
}

func Gimme(array [3]int) int {
	// присваиваем максимальное и минимальное число
	max := array[0]
	min := array[0]

	// ищем минимальное и максимальное
	for _, num := range array {
		if num > max {
			max = num
		}
		if num < min {
			min = num
		}
	}

	// ищем средний элемент
	for ind, num := range array {
		if num != max && num != min {
			return ind
		}
	}
	return -1
}
