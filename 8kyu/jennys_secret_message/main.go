package main

import "fmt"

func main() {
	fmt.Println(Greet("Alfred"))
	fmt.Println(Greet("Johnny"))
}

func Greet(name string) string {
	if name == "Johnny" {
		return fmt.Sprint("Hello, my love!")
	} else {
		return fmt.Sprintf("Hello, %v!", name)
	}
}
