package main

import "fmt"

func Music() []string {
	Music := []string{"Old town road", "Naanchaku", "Go to sleep", "No brainer", "2u"}
	Need := Music[:len(Music)-4]
	Out := make(([]string), len(Music))
	copy(Out, Need)
	return Out

}

func main() {
	Out := Music()
	fmt.Println(Out, len(Out))
}
