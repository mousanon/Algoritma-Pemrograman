package main

import "fmt"

const NMAX int = 100

type list [NMAX]int

func main() {
	var total, cari, i, k1, k2, b1, b2 int

	fmt.Scanln(&total, &cari)
	i = 1
	for i <= total {
		fmt.Scan(&tab[i])
		i = i + 1
	}

	if cari%2 == 0 {
		cariKecil(tab, total, &k1, &k2)
		fmt.Print("ditemukan")
	} else if cari%2 > 0 {
		cariBesar(tab, total, &b1, &b2)
		fmt.Print("tidak ditemukan")
	}
}

func cariKecil(T list, n int, k1 *int, k2 *int) {
	var i int

	*k1 = T[1]
	*k2 = 0
	i = 2
	for i < n {
		if *k1 > T[i] {
			*k1 = T[i]
			*k2 = *k1
		}
		i = i + 1
	}
}

func cariBesar(T list, n int, b1 *int, b2 *int) {\
	var i int

	*b1 = T[1]
	*b2 = 0
	i = 2
	for i < n {
		if *b1 < T[i] {
			*b1 = T[i]
			*b2 = *b1

		}
		i = i + 1
	}
}
