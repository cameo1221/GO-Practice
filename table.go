package main

import (
	"fmt"
)

func Calc(l, w float64) (area, perimeter float64) {
	area = l * w
	perimeter = 2 * (l + w)

	return area, perimeter
}

func main() {
	l := [...]float64{2.4, 5.5, 10, 22, 4}
	w := [...]float64{3.8, 7.9, 11, 25, 6}
	var ar []float64
	var pr []float64
	for i := 0; i < len(l); i++ {
		for j := 0; j < len(w); j++ {
			length := l[i]
			width := w[j]
			area, perimeter := Calc(length, width)
			ar = append(ar, area)
			pr = append(pr, perimeter)
			fmt.Printf("\n\nLength of the array 'ar' is : %v and capacity is : %v\n", len(ar), cap(ar))
			fmt.Println("\nLength of the array 'pr' is :", len(pr), "and capacity is :\n", cap(pr))

			fmt.Printf("\n%v %v Area %v of length %v and width %v ", i, j, area, length, width)
			fmt.Printf("\n%v %v Perimeter %v of length %v and width %v ", i, j, perimeter, length, width)
		}
	}
	fmt.Printf("\n\nThe total area is: %v\n", ar)
	fmt.Printf("\nThe total perimeters are: %v \n", pr)

}
