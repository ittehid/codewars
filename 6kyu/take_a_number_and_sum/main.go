package main

import (
	"fmt"
	"math/big"
	"strconv"
)

func main() {
	fmt.Println(SumDigPow(1, 10))
	fmt.Println(SumDigPow(10, 89))
	fmt.Println(SumDigPow(10, 100))
	fmt.Println(SumDigPow(90, 100))
	fmt.Println(SumDigPow(89, 135))
	fmt.Println(SumDigPow(12157692622039623047, 12157692622039625606))
}

func SumDigPow(a, b uint64) []uint64 {
	arr := make([]uint64, 0)
	for i := a; i <= b; i++ {
		str := strconv.FormatUint(i, 10)
		sum := big.NewInt(0)
		for j := 0; j < len(str); j++ {
			digit := int64(str[j] - '0')
			power := big.NewInt(0).Exp(big.NewInt(digit), big.NewInt(int64(j+1)), nil)
			sum.Add(sum, power)
		}
		if sum.Uint64() == i {
			arr = append(arr, i)
		}
	}
	return arr
}
