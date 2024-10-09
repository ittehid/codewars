package main

import "fmt"

func main() {
	fmt.Println(GetVolumeOfCuboid(1.0, 2.0, 2.0))
	fmt.Println(GetVolumeOfCuboid(6.3, 2.0, 5.0))
}

func GetVolumeOfCuboid(length, width, height float64) float64 {
	return length * width * height
}
