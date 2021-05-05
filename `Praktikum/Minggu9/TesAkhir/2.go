package main

import "fmt"

const NMAX = 1000000    // memberikan nilai konstan NMAX
var data [NMAX]int    // variabel data 

func main() {
	var n, k int
	fmt.Scan(&n, &k)

	isiArray(n)

	posisi := posisi(n, k)
	if posisi != -1 {
		fmt.Println(posisi)
	} else {
		fmt.Println("TIDAK ADA")
	}
}

func isArray(n int) {
	func isiArray(n int) {
		if n <= NMAX {
			for i := 0; i < n; i++ {
				fmt.Scan(&data[i])
			}
		}
	}
}

func posisi(n, k int) int {    // fungsi posisi
	for i := 0; i < n; i++ {
		if k[i] == x {
			return i
		}
	}
	return -1
}