package main

import "fmt"

const NMAX = 100    // mendeklarasikan NMAX

type mahasiswa struct {    // membuat tipe bentukan mahasiswa yang terdiri dari nama dan tinggi
	nama   string     // mendeklarasikan variabel nama dengan tipe data string
	tinggi int     // mendeklarasikan variabel tinggi dengan tipe data integer
}

type dataMhs [NMAX]mahasiswa    // mengkonversi tipe bentukan mahasiswa menjadi array dengan nama variabel dataMhs

func main() {    // program utama
	var n int    // mendeklarasikan variabel n dengan tipe data integer
	var t dataMhs    // mendeklarasikan varabel t degan tipe bentukan dataMhs
	fmt.Scanln(&n)    // input user disimpan ke n
	bacaData(&t, *n)    // memanggil bacaData
	urutData(&t, n)    // memanggil urutData
	cetakData(t, n)    // memanggil cetakData

}

func bacaData(t *dataMhs, n *int) {    // procedure bacaData
	/* IS. n data mahasiswa telah siap pada piranti masukan
	FS. menerima input n dan array t berisi n data tinggi mahasiswa */
	for i := 0; i < n; i++ {    // selama i kurang dari n lakukan program dibawah dan menambah 1 i setiap iterasi
		fmt.Scan(&t[i].nama, &t[i].tinggi)    // user input untuk array t dengan indeks i dengan value nama dan tinggi
	}
}
func cetakData(t dataMhs, n int) {     // procedure cetakData
	/* IS. terdefinisi sebuah array t yang berisi n data tinggi mahasiswa
	FS. menampilkan array t ke layar monitor */
	for i := 0; i < n; i++ {    // selama i kurang dari n dan melakukan iterasi maka lakukan program dibawah
		fmt.Print(t[i], " ")    // mencetak array t dengan indeks i kelayar
	}
}

func urutData(t *dataMhs, n int) {
	/* IS. terdefinisi sebuah array t yang berisi n data tinggi mahasiswa
	FS. array t terurut ascending berdasarkan tinggi dengan algoritma selection sort
	*/
	for i := 0; i < n; i++ {    // selama i kurang dari n dengan iterasi i lakukan program dibawah
		minIndex := i    // mendeklarasikan miniIndex = 1
		for j := i + 1; j < n; j++ {    // selama j kurang dari n
			if t[minIndex].tinggi > t[j].tinggi {    // jika tinggi miniindex lebih besar dari array j tinggi lakukan program dibawah
				minIndex = j    // mendeklarasikan miniindex = j
			}
		}
		temp := t[i].tinggi
		t[i].tinggi = t[minIndex].tinggi
		t[minIndex].tinggi = temp

		temp2 := t[i].nama
		t[i].nama = t[minIndex].nama
		t[minIndex].nama = temp2
	}
}
