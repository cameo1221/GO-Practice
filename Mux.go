package main

import "fmt"

func main() {
	for no, i := 10, 1; i < 10 && no < 20; i, no = i+1, no+1 {
		fmt.Printf("%d * %d = %d\n", no, i, no*i)
	}
}
