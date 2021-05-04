package main

import "fmt"

func hitungSkor(soal, skor *int){
	var x int

	for i := 0; i < 8; i++{
		fmt.Scan(&x)

		if x < 301 {
			*soal++
			*skor = *skor + x
		}
	}
}

func main() {
	var nama, nimi string
	var soal1, skor1, soal2, skor2 int
	
	fmt.Scan(&nama)
	hitung(&soal1, &skor1)

	fmt.Scan(&nimi)
	hitung(&soal2, &skor2)

	if soal1 > soal2 {
		fmt.Print(nama, soal1, skor1)
	} else {
		fmt.Print(nimi, soal2, skor2)
	}

}