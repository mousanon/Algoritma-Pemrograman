package main

import "fmt"

type arr = [20]int    // membuat array integer dengan range 20

func main() {
	var suara arr    // mendeklarasikan variabel suara dengan tipe data arr
	var suaraMasuk, suaraSah int    // mendeklarasikan suaraMasuk dan suaraSah dengan tipe data integer
	var i int = 0    // mendeklarasikan variabel i dengan tipe data integer dan memberikan nilai 0
	var x int    // mendeklarasikan variabel x dengan tipe data integer

	fmt.Scan(&x)    // variabel x menerima input dari user
	for x != 0 {    // selama x tidak sama dengan 0 lakukan program dibawah
		suaraMasuk++    // suaraMasuk bertambah satu
		if x > 0 && x <= 20 {    // jika x lebih besar dari 0 dan x lebih kecil sama dengan 0 maka lakukan program dibawah
			suaraSah++   // suaraSah bertambah satu
			suara[i] = x    // nilai array suara dengan index i dimasukkan ke dalam variabel x
			i++    // iterasi i
		}    // endif
		fmt.Scan(&x)    // variabel x menerima input dari user
	}    // endfor
	fmt.Println("Suara masuk: ", suaraMasuk)    // menarmpilkan ke layar jumlah suaraMasuk
	fmt.Println("Suara sah: ", suaraSah)    // menampilkan ke layar jumlah suaraSah

	var jumlah int    // mendeklarasikan variabel jumlah dengan tipe data integer
	for j := 1; j <= 20; j++ {   // jika nilai j kurang dari sama dengan 20 lakukan program dibawah
		jumlah = 0    // memeberikan nilai 0 pada jumlah
		
		ketuaWakil(suara, j, suaraSah, &jumlah)    // memanggil sub prosedur ketuaWakil ke dalam program utama

		if jumlah != 0 {    // jika jumlah tidak sama dengan 0 maka
			fmt.Println(j,": ", jumlah)    // menampilkan ke layar nilai j dan jumlah
		}
	}
}

func ketuaWakil(arraySuara arr, suara int, suaraSah int, jumlah *int) {    // prosedur ketuaWakil
	for i := 0; i < suaraSah; i++ {
		
	}
}