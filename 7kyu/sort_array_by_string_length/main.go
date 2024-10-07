package main

import (
	"fmt"
)

func main() {
	arr1 := []string{"beg", "life", "i", "to"}
	arr2 := []string{"", "moderately", "brains", "pizza"}
	fmt.Println(SortByLength(arr1))
	fmt.Println(SortByLength(arr2))
}

func SortByLength(arr []string) []string {
	// Проходим по всем элементам массива, кроме последнего
	for i := 0; i < len(arr)-1; i++ {
		// Предполагаем, что текущий индекс является индексом минимального элемента
		minIndex := i
		// Ищем минимальный элемент в оставшейся части массива
		for j := i + 1; j < len(arr); j++ {
			// Если длина строки на позиции j меньше, чем у текущего минимального
			if len(arr[j]) < len(arr[minIndex]) {
				// Обновляем индекс минимального элемента
				minIndex = j
			}
		}
		// Обмениваем местами текущий элемент с найденным минимальным
		arr[i], arr[minIndex] = arr[minIndex], arr[i]
	}
	// Возвращаем отсортированный массив
	return arr
}
