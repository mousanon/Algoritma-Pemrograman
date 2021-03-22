package main

import "fmt"

func Division(a, b int, result, remainder *int) {
	*result = a / b
	*remainder = a % b
}

func Num25tr(x int) {
	var a, b, result, remainder int
	Division(a, b, &result, &remainder)
	
}

func Des2Bin() {
	
}

func main() {
	var angka, a, b, c, d int
	fmt.Scanln(&angka)
	Division(a, b, &c, &d)
}