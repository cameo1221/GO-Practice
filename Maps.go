package main

import "fmt"

func main() {
	Movies := map[string]int{
		"Another round": 2019, "Inception": 2011,
		"Oppenheimer": 2023, "Babylon": 2023, "Bicyle thief": 1949,
	}
	for l, m := range Movies {
		fmt.Printf("\n %s was launched in %v", l, m)
	}
}
