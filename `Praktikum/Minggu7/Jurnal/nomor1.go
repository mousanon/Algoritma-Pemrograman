package main

import "fmt"

func main() {
	var n, p, q, kecil1, kecil2 int

	fmt.Scanln(&n, &p, &q)
	if p == q {
		fmt.Println("Kita tidak pernah bertemu")
	} else if n % p == 0 || n % q == 0 {
		fmt.Println("Kita bertemu hari ini")
	} else if n % p == 0 && n % q == 0 {
		if p > q {
			fmt.Println("Kita bertemu ", q , "hari lagi")
		} else {
			fmt.Println("Kita bertemu ", p , "hari lagi")
		}
	} else {
		kecil1 = p - (n % p) 
		kecil2 = q - (n % q)
		if kecil1 > kecil2 {
			fmt.Println("Kita bertemu ", kecil2 , " hari lagi")
		} else {
			fmt.Println("Kita bertemu ", kecil1 , " hari lagi")
		}
	}
}