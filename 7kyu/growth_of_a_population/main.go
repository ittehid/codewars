package main

import (
	"fmt"
	"math"
)

func main() {
	fmt.Println(NbYear(1500, 5, 100, 5000))
	fmt.Println(NbYear(1500000, 2.5, 10000, 2000000))
	fmt.Println(NbYear(1500000, 0.25, 1000, 2000000))
	fmt.Println(NbYear(1500000, 0.25, -1000, 2000000))
	fmt.Println(NbYear(1000, 1000, -1000, 2000000))
}

func NbYear(p0 int, percent float64, aug int, p int) int {
	years := 0
	p_current := float64(p0)
	actualPercent := percent / 100.0

	for p_current < float64(p) {
		p_current = p_current + (p_current * actualPercent) + float64(aug)
		p_current = math.Floor(p_current)
		years++
	}
	return years
}
