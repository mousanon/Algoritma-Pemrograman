package main

import "fmt"

func main() {
	var a, b int  // mendeklarasikan variabel a dan b merupakan integer

	count_pos := 0  // menetapkan nilai awal count_pos = 0 ; untuk menghitung hasil positif
	count_net := 0  // menetapkan nilai awal count_net = 0 ; untuk menghitung hasil netral
	// fmt.Scanln(&a, &b)
	for a + b >= 0 {  // selama pernyataan terpenuhi lakukan program dibawah
		fmt.Scanln(&a, &b)  // menginputkan a dan b
		if a + b == 0 {  // jika syarat terpenuhi lakukan program dibawah
			count_net++  // iterasi count_net
		}
		if a + b > 0 {  // jika syarat terpenuhi lakukan program dibawah
			count_pos++  // iterasi count_pos
		}
		if a + b < 0 {  // jika syarat terpenuhi lakukan program dibawah
			fmt.Println("Netral:", count_net)   // menampilkan ke layar jumlah hasil netral
			fmt.Println("Positive:", count_pos)  // menampilkan ke layar jumlah hasil positif
		}
	}
	
}