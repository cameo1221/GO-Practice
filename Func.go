package main

import "fmt"

func main() {
	Mood, name := getMood()
	fmt.Printf("My Name is %s and My mood is: %s", name, Mood)
}
func getMood() (string, string) {
	return "Sad", "Aryan"

}
