package main

import "fmt"

func BacaData(usaha *int, jumlahDoa *int, doaOrtu *bool, nilai *byte) {
	fmt.Print("Banyaknya usaha? ")
	fmt.Scanln(&*usaha)
	fmt.Print("Banyaknya doa? ")
	fmt.Scanln(&*jumlahDoa)
	fmt.Print("Doa orang tua? ")
	fmt.Scanln(&*doaOrtu)
	fmt.Print("Nilai Algoritma? ")
	fmt.Scanln(&*nilai)
}

func TabungUsahaDoa(usaha int, doa int, total *int) {
	*total = usaha + doa
}

func TabungDoaOrtu(doaOrtu bool, total int) int {
	if doaOrtu {
		return total * 2
	}
	return total
}

func HasilNilaiAlpro(nilai byte, total *int) {
	if nilai == 'A' {
		*total -= 150
	} else if nilai == 'B' {
		*total -= 130
	} else if nilai == 'C' {
		*total -= 100
	}
}

func HasilAkhir(poin int) string {
	if poin >= 130 {
		return "Lulus langsung dapat kerja gaji 2 digit"
	} else if poin >= 50 && poin < 130 {
		return "Lulus langsung dapat kerja"
	}
	return "Jangan lelah berdoa dan berusaha, tidak ada yang sia-sia dari usaha dan doa"
}

func main() {
	var total, usaha, nDoa int
	var doaOrtu bool
	var nilai byte

	BacaData(&usaha, &nDoa, &doaOrtu, &nilai)

	TabungUsahaDoa(usaha, nDoa, &total)
	total = TabungDoaOrtu(doaOrtu, total)
	HasilNilaiAlpro(nilai, &total)

	fmt.Println(HasilAkhir(total))
}