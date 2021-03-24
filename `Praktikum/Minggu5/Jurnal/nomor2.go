package main

import "fmt"

func f(x float64) float64 {    // function f
	return x*x    // me-return x*x
}

func g(x float64) float64 {    // fucntion g
	return x-2    // me-return x-2
}

func h(x float64) float64 {    // fucntion h
	return x+1    // me-return x+1
}

func komposisi(x float64, y *float64) {    // procedure komosisi
	*y = f(g(h(x)))    // menyimpan hsail perhitungan ke variabel y
}

func main() {
	var x, y float64    // mendeklarasikan variabel x dan y berupa integer
	fmt.Scanln(&x)    // menginputkan nilai ke variabel x

	komposisi(x,&y)    // memanggil procedure komposisi
	fmt.Println(y)    // menampilkan nilai y ke layar
}