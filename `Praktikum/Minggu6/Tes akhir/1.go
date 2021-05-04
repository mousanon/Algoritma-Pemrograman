package main

import (
	"fmt"
	"math"
)

func jarak(x1, y1, x2, y2 int) int {    // function jarak
	return sqrt(((x1-x2)(x1-x2)+(y1-y2)(y1-y2)))    // me-retrurn rumus jarak dari lingkaran
}

func lingkaran(c1, c2 int) int {    // function lingkaran
	if c1 < a1 && c2 < a2 && c1 < b1 && c2 < b2 {    // jika syarat if terpenuhi melakukan program dibawah
		return true   // jika if statement terpenuhi maka me-return true
	} else {    // jika if statement tidak terpenuhi, mengeksekusi else
		return false    //jika 
	}
}

func cariTeks() {
	dlmLuar := lingkaran(c1, c2)
	if dlmLuar == true {
		fmt.Println("di dalam lingkaran 1 dan 2")
	} else {
		fmt.Prinln("di luar lingkaran 1 dan 2")
	}
}

func main() {
	var a1, a2, a3 int
	var b1, b2, b3 int
	var c1, c2 int
	var dist int
	var dlmLuar bool

	fmt.Scanln(&a1, &a2, &a3)
	fmt.Scanln(&b1, &b2, &b3)
	fmt.Scanln(&c1, &c2)

	dist = jarak(a1, b1, a2, b2)
	dlmLuar = lingkaran(c1, c2)
	cariTeks()
}