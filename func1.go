package main

import "fmt"

func yearsuntill(Age int) (
	Adult int,
	Driving int,
	Drinking int,
) {
	Adult = 18 - Age
	if Adult < 1 {
		Adult = 0
	}
	Driving = 16 - Age
	if Driving < 1 {
		Driving = 0
	}
	Drinking = 22 - Age
	if Drinking < 1 {
		Drinking = 0
	}
	return
}

func test(Age int) {
	Adult, Driving, Drinking := yearsuntill(Age)
	fmt.Printf("Your Current Age is: %v", Age)
	for Age < 18 {
		Age += 1
		fmt.Println(Age)
	}

	if Age < 18 {
		fmt.Println("\nYou are a minor.")
	}

	fmt.Printf("\nYears untill Adult: %v", Adult)
	fmt.Printf("\nYears untill Driving: %v", Driving)
	fmt.Printf("\nYears untill Drinkig: %v", Drinking)
	fmt.Println("\n----------------------------------\n----------------------------------")
	return

}
func main() {
	test(10)
	test(17)
	test(22)
}
