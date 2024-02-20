package main

import "fmt"

func average(x []float64) (avg float64) {
	total := 0.0
	if len(x) == 0 {
		avg = 0
	} else {
		for _, v := range x {
			total += v
		}
		avg = total / float64(len(x))
	}
	return
}

func main() {
	x := []float64{12.12, 33.64, 12.102, 34.98}
	fmt.Println(average(x))
}
