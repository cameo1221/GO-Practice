package main

import (
	"fmt"
	"time"
)

func digit(number int, dgs chan int) {
	for number != 0 {
		digit := number % 10
		dgs <- digit
		number /= 10
	}
	close(dgs)

}

func Squareop(number int, sq chan int) {
	fmt.Println("Welcome to Node 1")
	sum := 0
	digits := make(chan int)
	go digit(number, digits)
	for digit := range digits {
		sum += digit * digit
	}
	sq <- sum
}

func cubeop(number int, cq chan int) {
	fmt.Println("Welcome to Node 2")
	sum := 0
	digits := make(chan int)
	go digit(number, digits)
	for digit := range digits {
		sum += digit * digit * digit
	}
	cq <- sum
}

func main() {
	start := time.Now()
	number := 9293
	Square := make(chan int)
	Cube := make(chan int)
	go Squareop(number, Square)
	go cubeop(number, Cube)
	squares, cubes := <-Square, <-Cube
	elapsed := time.Since(start)
	fmt.Printf("\nThe square of the digits is: %d \n", squares)
	fmt.Printf("The cube of the digits is: %d \n", cubes)
	fmt.Printf("\n THE TOTAL IS %d", squares+cubes)
	fmt.Printf("\nTime taken for execution :%s ", elapsed)
}
