package main

import "fmt"

type Describer interface {
	Describe()
}

type worker struct {
	name string
	age  int
}

func (w worker) Describe() {
	fmt.Printf("%s is %d years old", w.name, w.age)

}

func findTypes(i interface{}) {
	switch v := i.(type) {
	case Describer:
		v.Describe()
	default:
		fmt.Println("Error Banda margya")
	}
}

func main() {
	findTypes("Aman")
	p := worker{
		name: "Aryan",
		age:  21,
	}
	findTypes(p)
}
