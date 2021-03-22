package main

import "fmt"

func main() {
	var winner, player, temp string  // mendekarasikan winner, player dan temp sebagai string
	var nilai, answer int  // mendeklarasikan nilai dan answer sebagai integer

	winner = "A"  // menetapkan winner sebagai A
	player = "B"  // menteapkan player sebagai B
	ronde := 1  // menetapkan nilai awal ronde = 1
	fmt.Print("Ronde", ronde, ":")  // menampilkan ronde ke layar
	fmt.Print(winner, " - masukkan angka rahasia:")  // menampilkan tampilan untuk input
	fmt.Scanln(&nilai)  // menginputkan angka rahasia
	for nilai != -101 {  // selama nilai tidak sama dengan -101 lakukan program dibawah
		i := 1  // menetapkan nilai awal i = 1
		fmt.Print(player, "- masukkan angak tebakan ke-", i, ":") // menampilkan ke layar
		fmt.Scanln(&answer)  // menginputkan answer
		for i < 3 && answer != nilai { // selama i dan answer memenuhi lakukan program dibawah
			i++  // iterasi i
			fmt.Print(player, "- masukkan angka tebakan ke-", i, ":")  // menampilkan ke layar
			fmt.Scanln(&answer)  // menginputkan answer

		}
		if nilai == answer {  // jika nilai sama dengan answer lakukan program dibawah
			temp = winner  // temp sama dengan winner
			winner = temp  // winner sama dengan temp
			player = temp   // player sama dengan temp
		}
		fmt.Print(winner, " adalah pemenangnya\n")  // menampilkan pemenang ke layar
		ronde++  // iterasi ronde
		fmt.Print("Ronde", ronde, ":")  // menampilkan ronde ke layar
		fmt.Print(winner, "- masukkan angka rahasia:")  // menampilkan ke layar
		fmt.Scanln(&nilai)  // menginputkan nilai
		}
	fmt.Print("Permainan Selesai")  // menampilkan ke layar bahwa permainan selesai
}