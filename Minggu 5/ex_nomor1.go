package main

import "fmt"

func luasSegitiga(alas, tinggi float64, luas *float64) {
	fmt.Scan(&alas, &tinggi)
	*luas = (alas*tinggi)/2
}

func main() {
	var alas, tinggi, luas float64
	luasSegitiga(alas, tinggi, &luas)
	fmt.Print("Luas segitiga : ", luas)
}