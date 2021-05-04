package main

import "fmt"

func Kabisat(tahun int) bool {
	if tahun % 4 == 0 {
		return true
	} else {
		return false
	}
}

func Bulan2Angka(bulan string) int {
	if bulan == "januari" {
		return 1
	} else if bulan == "februari" {
		return 2
	} else if bulan == "maret" {
		return 3
	} else if bulan == "april" {
		return 4
	} else if bulan == "mei" {
		return 5
	} else if bulan == "juni" {
		return 6
	} else if bulan == "juli" {
		return 7
	} else if bulan == "agustus" {
		return 8
	} else if bulan == "september" {
		return 9
	} else if bulan == "oktober" {
		return 10
	} else if bulan == "november" {
		return 11
	} else if bulan == "desember" {
		return 12
	}
}

func JumlahHari(bln, thn int) int {

}

func main() {
	var tanggal, tahun int
	var bulan, hari string
	
	fmt.Scanln(&tanggal, &bulan, &tahun, &hari)
	Kabisat(tahun)
	Bulan2Angka(bulan)
	JumlahHari(bulan, tahun)
}