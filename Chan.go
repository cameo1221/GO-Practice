package main

import (
	"fmt"
)

func Rt(done chan bool) {
	fmt.Printf("Hello This is Node 1")
	done <- true
}

func main() {
	done := make(chan bool)
	go Rt(done)
	<-done

	fmt.Println("\n end of node 2")
}
