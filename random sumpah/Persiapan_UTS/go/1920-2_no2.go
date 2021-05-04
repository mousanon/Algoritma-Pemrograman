package main

import "fmt"

func main() {
	var a, b, x int

	fmt.Print("a,b")
	fmt.Scanln(&a, &b)
	for a; a < b {
		b = a + 1
		x = a + b

	}
	fmt.Println("nilai x = ", x)
}