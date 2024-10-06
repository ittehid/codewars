package main

import "fmt"

func main() {
	items := []interface{}{1, 2, 3, 4, 5, 6, 7}
	k := 3

	fmt.Println(Josephus(items, k))
}

func Josephus(items []interface{}, k int) []interface{} {
	result := make([]interface{}, 0)
	index := 0
	for len(items) > 0 {
		index = (index + k - 1) % len(items)
		result = append(result, items[index])
		items = append(items[:index], items[index+1:]...)
	}
	return result
}
