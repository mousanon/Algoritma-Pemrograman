package main

import "fmt"

func main() {
	var n int   // mendeklarasikan n sebagai nteger
	var mi, mo, ma int  // mendeklarasikan mi, mo dan ma sebagai nteger

	jum_ma := 0  // menetapkan nilai awal jum_ma = 0
	i := 1  // menetapkan nilai awal i = 0
	tahun := 0  // menetapkan nilai awal tahun = 0
	fmt.Scanln(&n)  // menginputkan n
	for i <= n {  // selama i lebih kurang atau sama dengan n lakukan program dibawah
		fmt.Scanln(&mi, &mo, &ma) // menginputkan mi, mo, dan ma
		if jum_ma + mi - mo != ma && tahun == 0 {  // jika if statement benar lakukan program dibaawah
			tahun = i   // menetapkan tahun sama dengan i
		}
		jum_ma = ma  // menetapkan jum_ma sama dengan ma
		i++  // iterasi dari perulangan
	}
	fmt.Println(tahun)  // menampilkan ke layar variabel tahun
}