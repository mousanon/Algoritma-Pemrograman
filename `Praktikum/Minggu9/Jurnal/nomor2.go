package main

import "fmt"

type arr [100]string

func main() {
	var n, m int

	var mahasiswa, wisuda arr

	fmt.Scan(&n)
	for i := 0; i < n; i++ {
		fmt.Scan(&mahasiswa[i])
	}

	fmt.Scan(&m)
	for i := 0; i < m; i++ {
		fmt.Scan(&wisuda[i])
	}

	updateKelulusan(&mahasiswa, wisuda, &n, m)

	fmt.Print(n, " ")
	for i := 0; i < n; i++ {
		fmt.Print(mahasiswa[i], " ")
	}
}

func posisi(tab arr, n int, x string) int {
	for i := 0; i < n; i++ {
		if tab[i] == x {
			return i
		}
	}
	return -1
}

func hapus(tab *arr, n *int, x string) {
	data := posisi(*tab, *n, x)
	if data != -1 {
		for i := data; i < *n; i++ {
			tab[i] = tab[i+1]
		}
		*n--
	}
}

func updateKelulusan(mhs *arr, wisuda arr, n *int, m int) {
	for i := 0; i < m; i++ {
		hapus(mhs, n, wisuda[i])
	}
}