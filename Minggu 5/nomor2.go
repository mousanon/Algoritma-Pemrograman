package main

import "fmt"

func isTriangle(a, b, c int) {
	if a+b <= c || a+c <= b || b+c <= a {
		fmt.Println("Buksn Segitiga")
	} else {
		fmt.Println("Segitiga")
	}
}

func main() {
	var a, b, c int
	fmt.Scanln(&a, &b, &c)
	isTriangle(a, b, c)
}