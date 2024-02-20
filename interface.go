package main

import "fmt"

type VowelsCounter interface {
	Findvowels() []rune
}

type Mystring string

func (ms Mystring) Findvowels() []rune {
	var vowels []rune
	var count int
	for _, rune := range ms {
		if rune == 'a' || rune == 'e' || rune == 'i' || rune == 'o' || rune == 'u' {
			vowels = append(vowels, rune)
			count += 1
		}
	}
	fmt.Println("Number of vowels are:", count)
	return vowels
}

func main() {
	name := Mystring("aeiou")
	var v VowelsCounter
	v = name
	fmt.Printf("Vowels in the given string : %c ", v.Findvowels())

}
