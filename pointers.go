package main

import "fmt"

func main() {
	var address *int
	number := 69
	address = &number
	value := *address
	fmt.Printf("The address is %v and the value is %v", address, value)
}
