package main

import (
	"fmt"
	"strconv"
)

func main() {
	fmt.Println(BinToDec("0"))
	fmt.Println(BinToDec("1"))
	fmt.Println(BinToDec("10"))
	fmt.Println(BinToDec("11"))
	fmt.Println(BinToDec("101010"))
	fmt.Println(BinToDec("1001001"))
}

func BinToDec(bin string) int {
	result, _ := strconv.ParseInt(bin, 2, 0)
	return int(result)
}
