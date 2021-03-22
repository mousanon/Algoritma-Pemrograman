package main
import "fmt"

func main() {
	var digit, angka int   // mendeklarasikan digit dan anggka dalam bentuk integer
	i := 0   // mendeklarasikan variabel i adalah 0
	hasil := 0   // mendeklarasikan variabel hasil adalah 0

	fmt.Scanln(&digit) // menginputkan digit 
	for i<digit {    // selama i didalam jangkauan digit jalankan kode dibawahnya
		fmt.Scanln(&angka)   // inputkan angka
		if x >= 0 && x <= 9 {   // jika x lebih besar sama dengan 0 dan x kurang dari sama dengan 9 jalankan kode dibawahnya
			j := 0   // mendeklarasikan j sama dengan 0
			for j<i; j++ {   // selama j lebih kecil dari i jalankan kode dibawahnya
				x = x*10    // jika memenuhi kode diatas maka x dikalikan dengan 10
			}
				hasil = hasil + x   // hasil ditambah dengan x
				i++   // iterasi i
		}
	}
	fmt.Println(hasil)   // tampilkan hasil ke layar
}