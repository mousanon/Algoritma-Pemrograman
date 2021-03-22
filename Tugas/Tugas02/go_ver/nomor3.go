package main 

import "fmt"

func BacaData(usaha int, jumlahDoa int, doaOrtu bool, nilai string) {
	fmt.Print("Banyaknya usaha? ")
	fmt.Scanln(&usaha)
	fmt.Println(usaha)
	fmt.Print("Banyaknya doa? ")
	fmt.Scanln(&jumlahDoa)
	fmt.Print("Doa orang tua? ")
	fmt.Scanln(&doaOrtu)
	fmt.Print("Nilai Algoritma? ")
	fmt.Scanln(&nilai)
	
}

func TabungUsahaDoa(usaha int, doa int, total int) {
	
}

// func TabungDoaOrtu(doaOrtu bool, total int) int {
// 	if doaOrtu == true {
// 		total = total * 2
// 	}
// 	return total
// }

func main() {
	var usaha, jumlahDoa, total int
	var doaOrtu bool
	var nilai string
	BacaData(usaha, jumlahDoa, doaOrtu, nilai)

	TabungUsahaDoa(usaha, jumlahDoa, total)
	fmt.Print(total)
	// TabungDoaOrtu(doaOrtu, total)
}