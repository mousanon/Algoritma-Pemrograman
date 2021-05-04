package main

import "fmt"

func Division(a, b int, result, remainder *int) {    // procedure Division
	*result = a / b    // mengisi variabel result dengan membagi a dan b
	*remainder = a % b    // mengisi variabel remainder dengan a modulo b
}

func Num2Str(x int) string {    // function Num2Str
	if x == 0 {    // jika syarat terpenuhi, melakukan program dibawah
		return "0"     // jika if statement benar maka akan mereturn string 0
	} else if x == 1 {    // jika if statement tidak terpenuhi lakukan program dibawah
		return "1"    // jika else if benar maka akan mereturn string 1
	} else {    // jika kedua syarat diatas tidak terpenuhi melakukan program dibawah
		return ""    // jika else terpenuhi return string kosong
	}
}

func Des2Bin(desimal int) string {    // function Des2Bin
	var hasil string    // mendeklarasikan variabel lokal yang bernilai string
	var rDiv, rMod int    // mendeklarasikan rDiv dan rMod bernilai integer
	
	hasil = ""    // memberi nilai awal variabel hasil dengan string kosong
	for desimal > 0 {    // selama nilai desimal lebih dari 0 lakukan program dibawah
		Division(desimal, 2, &rDiv, &rMod)    // memanggil procedure Division
		desimal = rDiv    // mengisi nilai desimal dengan rDiv
		hasil = Num2Str(rMod) + hasil    // hasil berisi function Num2Str ditambah dnegan hasil
	}
	return hasil
}

func main() {    // program utama dari number
	var biner string     // mendeklarasikan variabel biner sebagai string
	var desimal int    // mendeklarasikan variabel desimal sebagai integer
	fmt.Scanln(&desimal)     // menginputkan integer ke desimal
	biner = Des2Bin(desimal)     // mengisi variabel biner dengan function Des2Bin
	fmt.Println(biner)     // menampilkan variabel biner ke layar
}