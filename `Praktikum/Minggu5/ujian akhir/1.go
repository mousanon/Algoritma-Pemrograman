package main 

import "fmt"

func minmax(f1, f2, f3, f4 int, min, max *int) {    // procedure minmax
	min = f1    // 
	max = f1

	if f2 > max {
		max = f2
	} else if f3 > max {
		max = f3
	} else if f4 > max {
		max = f4
	}

	if f2 < min {
		min = f2
	} else if f3 < min {
		min = f3
	} else if f4 < min {
		min = f4
	}
	fmt.Println(max, min)
}

func main() {
	var a, b, c, d int
	var max, min int
	fmt.Scanln(&a, &b, &c, &d)

	minmax(a, b, c, d, &min, &max)
}