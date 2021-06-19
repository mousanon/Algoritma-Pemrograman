package main

import "fmt"

const NMAX = 1000000

type partai struct {
	num, jumlah int
}	

type arrayPartai [NMAX]partai

func main() {
	var arr arrayPartai
	var inputan, cek int
	var counter int = 0

	fmt.Scan(&inputan)

	for inputan != -1 {
		cek = posisi(arr, counter, inputan)
		// panggil fungsi posisi
		if cek != -1 {
			arr[cek].jumlah++
		} else {
			arr[counter].num = inputan
			arr[counter].jumlah = 1
			counter++
		}
		fmt.Scan(&inputan)
	}
	for i := 0; i < counter; i++ {
		for j := i; j < counter; j++ {

		}
	}
	for i := 0; i < counter; i++ {
		fmt.Printf("%d(%d) ", arr[i].num, arr[i].jumlah)
	}
}

func posisi(arr arrayPartai, n int, dicari int) int {
	for i := 0; i < n; i++ {
		if arr[i].num == dicari {
			return i
		}
	}
	return -1
}
