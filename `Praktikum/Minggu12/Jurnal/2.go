package main

import "fmt"

func main() {
	var tabInt [100]int    // mendeklaraikan array integer dengan maksimal 100
	var n, d, u int    // mendeklarasikan variabel n, d, u dengan tipe data integer
	fmt.Scan(&n, &d, &u)    // memasukan integer kedalam n, d, dan u
	isiArray(&tabInt, &n)    // memanggil procedure isiArray
	sorting(&tabInt, u, d, n)    // memanggil procedure sorting
	tampilkanArray(tabInt, n)    // memanggil procedure tampilkanArray
}
func isiArray(t *[100]int, n *int) {    // procedure isiArray
	/* IS. terdefinisi bilangan bulat n, dan n buah bilangn bulat telah siap pada
	FS. array t berisi n buah bilangan yang berasal dari user
	*/
	for i := 0; i < *n; i++ {    // iterasi selama i < n menjalankan program dibawah
		fmt.Scan(&t[i])    // meminta user input ke dalam t dengan indeks i
	}
}

func tampilkanArray(t [100]int, n int) {    // procedure tampilkanArray
	/*
		IS. terdefinisi sebuah array t yang berisi n buah bilangan bulat
		FS. menampilkan array t ke layar secara horizontal
	*/
	for i := 0; i < n; i++ {    // selama i < n lakukan program dibawah
		fmt.Print(t[i]," ")    // mencetak array t dengan indeks i ke layar
	}
}

func sorting(t *[100]int, u, d, n int) {    // procedure sorting
	/*
		IS. terdefinisi sebuah array t yang berisi n bilangan bulat, dan indeks bilangan
		bulat u dan d
		FS. array t terurut descending dari indeks ke-u hingga ke-d dengan menggunakan algoritma insertion sort
	*/

	for i := u; i < d; i++ {    // selama i < d
		for j := i; j > 0 {    // selama j lebih dari 0 lakukan program dibawah
			if t[j-1] > t[j] {    // jika array t dengan indeks j-1 lakukan program dibawah
				t[j-1], t[j] = t[j], t[j-1]
			}
			j--    // iterasi pengurangan j
		}
	}
}