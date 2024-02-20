package main

import (
	"fmt"
	"time"
)

func Routine() {
	fmt.Println("Hello Go Routines!")
	return
}

func main() {
	go Routine()

	fmt.Println("Ma ka bhosda")
	time.Sleep(time.Second)
}
