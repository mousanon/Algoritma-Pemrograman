package main

import "fmt"

func b(x int) int {    // function b
	return (2*x) + 5    // me-return (2*x) + 5
}

func c(x int) int {    // function c
	return (15*x) + 30    // me-return (15*x) + 30
}

func d(x int) int {    // function d
	return x - 3    // mereturn x -3
}

func matematika(x int, y *int) {    // procedure matematika
	*y = d(c(b(x)))    // y menampung hasil perhitungan matematika
}

func main() {    // program utama
	var x int   // mendeklarasikan variabel x dengan tipe data x
	var y int    // mendeklarasikan variabel y dengan tipe data y
	var b, c, d string    // mendeklarasikan variabel b, c, d dengan itpe data string

	fmt.Scanf("%d %s %s %s", &x, &b, &c, &d)    // menginputkan x, b, c, d

	matematika(x, &y)    // memanggil procedure matematik
	fmt.Printf("%d adalah hasil (%s o %s o %s)(%d)\n", y, d, c, b, x)     // menampilkan ke layar hasil yang sudah dihitung
}