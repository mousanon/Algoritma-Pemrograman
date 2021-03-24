package main

import "fmt"

func findFactorial(n int, result *int) {    // procedure findFactorial
	var x int = 1    // mendeklarasikan x sebagai integer dab diberikan nilai 1
	for i := 1; i <= n; i++ {	// selama i memenuhi syarat lekukan program dibawah
		x = x * i    // kalikan x dengan i
	}
	*result = x    // result menampung nilai x
}

func permutation(n, r int) int {    // function permutation
	var f1, f2 int   // mendeklarasikan variabel lokal f1 dan f2 sebagai integer
	findFactorial(n, &f1)    // memanggil procedure findFactorial
	findFactorial(n-r,  &f2)    // memanggil procedure findFactorial

	return f1 / f2    // me-return hasil f1 / f2
}

func combination(n, r int) int {    // function combination
	var f1, f2, f3 int    // mendeklarasikan variabel lokal f1, f2 dan f3 sebagai integer
	findFactorial(n, &f1)    // memanggil findFactiorial
	findFactorial(n, &f2)    // memanggil findFactiorial
	findFactorial(n-r, &f3)    // memanggil findFactiorial

	return f1 / (f2 * f3)    // me-return hasil f1 / (f2 * f3)
}

func main() {    // program utama
	var n, r, p, c int    // mendeklarasikan variabel n, r, p dan c sebagai integer
	fmt.Scanln(&n, &r)    // menginputkan integer ke n dan r

	p = permutation(n, r)    // memangiil funciton permutation
	c = combination(n, r)    // memanggil function combination
	fmt.Println(p, c)    // menampilkan permutation dan combination ke layar
}