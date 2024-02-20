package main

import "fmt"

func main() {
	curr := "y"
	prev := "green"
	switch {
	case curr == "y" && prev == "r":
		fmt.Printf("Ready to move!")
	case curr == "y" && prev == "g":
		fmt.Printf("Wait!")
	default:
		fmt.Printf("Error!")
	}
}
