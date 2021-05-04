package main

import "fmt"

func hitungNopol(noPol string, bdg, jkt, lainnya *int) {
	if nopol == "B" || nopol == "F" {
		*jkt++
	} else if nopol == "D" || nopol == "Z" {
		*bdg++
	} else {
		*lainnya++
	}
}

var nopol string
s
func main() {
	var nBandung, nJakarta, lain int

	fmt.Scan(&nopol) 
	for nopol != "."{
		hitungNopol(nopol, nBandung, nJakarta, lain)
		fmt.Scan(&nopol)
	}
	fmt.Println(nBandung, nJakarta, lain)
}