package main

import "fmt"

func main() {
	fmt.Println(HumanReadableTime(0))
	fmt.Println(HumanReadableTime(59))
	fmt.Println(HumanReadableTime(60))
	fmt.Println(HumanReadableTime(3600))
	fmt.Println(HumanReadableTime(359999))
}

func HumanReadableTime(seconds int) string {
	hours := seconds / 3600
	minutes := (seconds % 3600) / 60
	seconds = seconds % 60
	return fmt.Sprintf("%02d:%02d:%02d", hours, minutes, seconds)
}
