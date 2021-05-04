package main

import "fmt"

func USDtoIDr(budget int) int {
	return budget * 13000
}

func hitungBiayaSatuan(inap, tempatWisata int, tiketMasuk, penginapan, konsumsi *int) {
	*tiketMasuk = tempatWisata * 13
	*penginapan = inap * 514
	*konsumsi = (inap * 2) * 20
}

func hitungTotalBiaya(tiketMasuk, pengiapan, konsumsi int, biaya *int) {
	*biaya = tiketMasuk + penginapan + konsumsi
}

func main() {
	var budget, tempatWisata, inap int
	var tiketMasuk, pengiapan, konsumsi int
	
	fmt.Print("Masukkan budget dalam dolar: ")
	fmt.Scanln(&budget)
	fmt.Print("Masukkan jumlah tempat wisata yang dikunjungi: ")
	fmt.Scanln(&tempatWisata)
	fmt.Print("Lama menginap: ")
	fmt.Scanln(&inap)
	
	USDtoIDr(budget)
	hitungBiayaSatuan(inap, tempatWisata, &tiketMasuk, &penginapan, &konsumsi)
	hitungTotalBiaya(tiketMasuk, penginapan, konsumsi, &biaya)
}