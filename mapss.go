package main

import "fmt"

type movie struct {
	genre  string
	rating int
}

func main() {
	m1 := movie{
		genre:  "comedy",
		rating: 7,
	}
	m2 := movie{
		genre:  "dark comedy",
		rating: 10,
	}
	m3 := movie{
		genre:  "teen comedy",
		rating: 5,
	}
	m4 := movie{
		genre:  "romantic comedy",
		rating: 8,
	}

	Movieinfo := map[string]movie{
		"Hera pheri": m1,
		"Game time":  m2,
		"superbad":   m3,
		"juno":       m4,
	}

	for name, info := range Movieinfo {
		fmt.Printf("\nName of the movie : %s\n Genere of the movie : %s \n Rating of the movie: %d\n length of map: %d\n", name, info.genre, info.rating, len(Movieinfo))
		delete(Movieinfo, "Hera pheri") // Deleting the entry for "Hera Pheri" from the

	}
}
