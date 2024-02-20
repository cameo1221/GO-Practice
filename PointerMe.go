package main

import "fmt"

type Age struct {
	age   int
	limit int
}

func Verify(a *Age) {
	if a.age < a.limit {
		fmt.Printf("Sorry you are not eligible for voting. You can apply in %d years.\n", (a.limit - a.age))
	} else if a.age > a.limit {
		fmt.Printf("You are older than the allowed age to vote. its been %d years since you voted\n", (a.age - a.limit))
	}
}

func (a *Age) Verify() {
	if a.age < a.limit {
		fmt.Printf("Sorry you are not eligible for voting. You can apply in %d years.\n", (a.limit - a.age))
	} else if a.age > a.limit {
		fmt.Printf("You are older than the allowed age to vote. its been %d years since you voted\n", (a.age - a.limit))
	}
}

func main() {
	a := Age{
		age:   15,
		limit: 18,
	}
	p := &a
	Verify(p)
	(p).Verify()
	(&a).Verify()
}
