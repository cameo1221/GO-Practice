package main

import "fmt"

func main() {
	const cost = 20
	Number := 4
	Number = amount(cost, Number)
	fmt.Println("The total price is :", Number)

}

func amount(cost, Number int) int {
	Number = cost * Number
	return Number
}
