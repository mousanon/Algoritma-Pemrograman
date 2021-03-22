package main

import "fmt"

func konversiWaktu(jam, menit, detik int) {
	fmt.Println(jam*3600 + menit*60 + detik)
}

func main() {
	var jam, menit, detik int
	fmt.Scanln(&jam, &menit, &detik)
	konversiWaktu(jam, menit, detik)
}