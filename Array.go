package main

import (
	"fmt"
	"reflect"
)

func main() {

	var movies = [3]string{"Joyland", "Out in Wild", "Babylon"}

	for i := 0; i < len(movies); i++ {
		options := movies[i]
		fmt.Println(i, options)
	}

	directors := [...]string{"Wes Anderson", " David Finch", "Sorcecce", "Terantino"}
	for index, option := range directors {
		fmt.Printf("%v: %s\n", index, option)
	}

	genres := [...]string{
		"Dark comedy", "Raunchy", "Thriller", "Biopic",
		"Inspiring", "Novel", "Art", "Theatrical",
	}
	Good_Mood := genres[0:2]
	Motivation := make([]string, 3)
	Motivation = genres[3:6]
	Good_Watch := genres[6:8]

	fmt.Printf("Good_Mood: %v \n Motivation: %v \n Good_Watch: %v", Good_Mood, Motivation, Good_Watch)
	fmt.Println(reflect.TypeOf(Good_Mood).Kind())

	all_Movies := append(genres[:], movies[:]...)

	fmt.Println(all_Movies)

	watch := all_Movies[4:8:9]
	fmt.Println(watch)

}
