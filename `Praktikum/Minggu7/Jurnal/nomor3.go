package main

import "fmt"

func main() {
	var a, b, besar int

	fmt.Scanln(&a, &b)
	if a > b {
		besar = a
	} else if b > a {
		besar = b
	}

	for besar % a != 0 || besar % b != 0 {
		besar++
	}
	fmt.Println(besar)
}