package main

import "fmt"

func main() {
	var N int
	var bb, rata float64

	i := 0
	fmt.Scanln(&N)
	for i < N {
		fmt.Scan(&bb)
		rata = bb / float64(N)
		i++
	}
	fmt.Println(rata)
}