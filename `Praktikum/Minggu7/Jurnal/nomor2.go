package main

import "fmt"

func hitung(suara, a, b, c *int) {
	if suara == 1 {
		*a++
	} else if suara == 2 {
		*b++
	} else if suara == 3 {
		*c++
	}
}

var suara int

func main() {
	var suaraA, suaraB, suaraC int

	fmt.Scan(&suara)
	for suara != -1 {
		hitung(suara, &suaraA, &suaraB, &suaraC)
		fmt.Scan(&suara)
	}
	fmt.Println(suaraA, suaraB, suaraC)
}