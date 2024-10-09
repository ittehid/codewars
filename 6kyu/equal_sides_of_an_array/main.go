package main

import "fmt"

func main() {
	arr1 := []int{1, 2, 3, 4, 3, 2, 1}
	arr2 := []int{-15, 5, 11, 17, 19, -17, 20, -6, 17, -17, 19, 16, -15, -6, 20, 17}
	arr3 := []int{-1, -2, -3, -4, -3, -2, -1}
	arr4 := []int{10, -10}

	fmt.Println(FindEvenIndex(arr1))
	fmt.Println(FindEvenIndex(arr2))
	fmt.Println(FindEvenIndex(arr3))
	fmt.Println(FindEvenIndex(arr4))
}

func FindEvenIndex(arr []int) int {
	arrSum := 0
	leftSum := 0

	// Общая сумма элементов массива
	for _, num := range arr {
		arrSum += num
	}

	// Проходим по массиву, обновляя суммы слева и справа от текущего индекса
	for i, num := range arr {
		arrSum -= num // Вычитаем текущий элемент из правой суммы
		if leftSum == arrSum {
			// Если суммы слева и справа равны, возвращаем текущий индекс
			return i
		}
		leftSum += num // Добавляем текущий элемент к левой сумме
	}
	// Если такой индекс не найден, возвращаем -1
	return -1
}
