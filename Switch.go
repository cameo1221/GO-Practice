package main

import "fmt"

func average(x []float64) (avg float64) {
	total := 0.00
	switch len(x) {
	case 0:
		avg = 0
	default:
		for _, v := range x {
			total += v

		}
		avg = total / float64(len(x))
	}
	return
}

func main() {
	x := []float64{12.23, 19.56, 55.43, 76.89}
	fmt.Println(average(x))
}
