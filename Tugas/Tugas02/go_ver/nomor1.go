package main

import "fmt"

func hitungRataRata(b float64, i int, r *float64) {
	*r = b / float64(i)
}

func main() {
	var bilangan, rata float64
	var i int

	fmt.Scanln(&bilangan)
	i = 0
	rata = 0
	for bilangan > 0 {
		i++
		hitungRataRata(bilangan, i, &rata)
		fmt.Scanln(&bilangan)
	}
	print(rata)
}