package main

import (
	"fmt"
	"unsafe"
)

func main() {
	a := 89
	b := 95

	fmt.Printf("The value of a %d and the value of b is %d", a, b)
	fmt.Printf("\nThe type of a is %T and the size of a is %d", a, unsafe.Sizeof(a))
	fmt.Printf("\nThe type of b is %T and the size of b is %d", b, unsafe.Sizeof(b))

}
