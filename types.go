package main

import "fmt"

func main() {
	//User specified
	const a int32 = 22
	const b string = "Arian"
	var c int64
	var d float32 = 1.223
	var e bool

	n := "Aryan"
	m := true
	l := 778.12

	fmt.Printf("Types of specified by user\n %T %T %T %T %T\n", a, b, c, d, e)
	fmt.Printf("Value of above\n %v %s %v %f %t\n", a, b, c, d, e)
	fmt.Printf("Type and value\n %T %T %T\n, \n %s %t %f \n", n, m, l, n, m, l)

}
