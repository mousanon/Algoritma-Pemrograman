package main

import "fmt"

func main() {
	var n, mi, mo, ma int
	fmt.Scanln(&n)
	for i := 0; i <= n; i++ {
		fmt.Scanln(&mi, &mo, &ma)
	}
}