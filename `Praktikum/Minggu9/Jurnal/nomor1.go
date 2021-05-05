package main

import "fmt"

type arr = [20]int

func main() {
	var suara arr 
	var suaraMasuk, suaraSah int
	var i int = 0	//index array
	var x int

	fmt.Scan(&x)
	for x != 0 {
		suaraMasuk++
		if x > 0 && x <= 20 {
			suaraSah++
			suara[i] = x
			i++
		}
		fmt.Scan(&x)
	}
	fmt.Println("Suara masuk: ", suaraMasuk)
	fmt.Println("Suara sah: ", suaraSah)

	var jumlah int
	for j := 1; j <= 20; j++ {
		jumlah = 0

		// prosedur searching
		searching(suara, j, suaraSah, &jumlah)

		if jumlah != 0 {
			fmt.Println(j,": ", jumlah)
		}
	}
}

func searching(arraySuara arr, suara int, suaraSah int, jumlah *int) {
	for i := 0; i < suaraSah; i++ {
		if arraySuara[i] == suara {
			*jumlah++
		}
	}
}