package main

import "fmt"


func Kabisat(tahun int) bool {
	return tahun % 400 == 0 || tahun % 4 == 0 && tahun % 100 != 0
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
	return 0
}

func Angka2Bulan(bulan int) string {
	if bulan == 1 {
		return "januari"
	} else if bulan == 2 {
		return "februari"
	} else if bulan == 3 {
		return "maret"
	} else if bulan == 4 {
		return "april"
	} else if bulan == 5 {
		return "mei"
	} else if bulan == 6 {
		return "juni"
	} else if bulan == 7 {
		return "juli"
	} else if bulan == 8 {
		return "agustus"
	} else if bulan == 9 {
		return "september"
	} else if bulan == 10 {
		return "oktober"
	} else if bulan == 11 {
		return "november"
	} else if bulan == 12 {
		return "desember"
	}
	return ""
}

func JumlahHari(bln int, thn int) int {
	if bln == 1 || bln == 3 || bln == 5 || bln == 7 || bln == 8 || bln == 10 || bln == 12 {
		return 31
	} else {
		if bln == 2 {
			if Kabisat(thn) {
				return 29
			}
			return 28
		}
		return 30
	}
}

func Pengambilan(tgl1 int, bln1 int, thn1 int, hari string, tgl2 *int, bln2 *int, thn2 *int) {
	var penambahanTanggal, jumlahHari int
	penambahanTanggal = 2
	jumlahHari = JumlahHari(bln1, thn1)
	if hari == "kamis" || hari == "jumat" {
		penambahanTanggal += 2
	}

	*tgl2 = tgl1 + penambahanTanggal
	*bln2 = bln1
	*thn2 = thn1

	if *tgl2 > jumlahHari {
		*tgl2 -= jumlahHari
		*bln2 = bln1 + 1
		if bln1 == 12 {
			*bln2 = 1
			*thn2 = thn1 + 1
		}
	}
}

func main() {
	var day1, day2, year1, year2, m2 int
	var day, month1, month2 string

	fmt.Scanln(&day1, &month1, &year1, &day)

	Pengambilan(day1, Bulan2Angka(month1), year1, day, &day2, &m2, &year2)
	month2 = Angka2Bulan(m2)

	fmt.Println(day2, month2, year2)
}