package main

import "fmt"

func main() {
	age := 29
	height := 174.62
	msg := fmt.Sprintf("Your age is: %v and height is %.1f", age, height)
	fmt.Println(msg)
}
