package main

import "fmt"

func main() {
	Empsalary := map[string]int{
		"Aryan": 2000,
		"Rahul": 1500,
		"Rah":   3500,
	}
	new := "joe"
	value, ok := Empsalary[new]
	if ok == true {
		fmt.Printf("The salary of %s is %d\n", new, value)

		return
	}
	fmt.Println("Employee not found")
}
