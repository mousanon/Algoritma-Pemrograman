package main

import "fmt"

const NMAX = 100

type belanja [NMAX]int

func main() {
	// Deklarasi Variabel
	var t belanja  // mendeklarasikan t merupakan array dari belanja
	var n, ha int  // mendeklaraiksan n dan ha merupakan integer
	var hp float64 // mendeklarasikan hp menjadi float
	// Lakukan proses masukan
	fmt.Scanln(&hp, &n) // menginputkan hp dan n
	// Hitung total belanja
	ha = total(t, n)
	// tentukan apakah mendapatkan promo atau tidak
	if ha > 100000 {
		// lakukan pengurutan
		urut(&t, n)
		// lakukan perhitungan promo
		promo(t, n, &hp)
		// tampilkan keluaran yang diminta
		fmt.Println(ha, int(hp))
	} else {
		// tampilkan keluaran yang diminta
		fmt.Println(ha, ha)
	}
}
func entri(t *belanja, n *int) {
	/* IS. data belanja telah siap pada piranti masukan
	FS. array t berisi sejumlah n harga barang yang dibeli */
	var harga, jumlah int
	fmt.Scan(&harga, &jumlah)
	for harga != 0 && jumlah != 0 {
		t[*n] = harga * jumlah
		*n++
	}
}
func urut(t *belanja, n int) {
	/* IS. terdefinisi array t yang berisi n harga barang yang dibeli
	FS. array t terurut secara ascending berdasarkan harga barang dengan
	selection/insertion sort */
	for i := 0; i < n; i++ {
		for j := i; j < n; j++ {
			if t[j] < t[i] {
				temp = t[i]
				t[i] = t[j]
				t[j] = temp
			}
		}
	}
}
func total(t belanja, n int) int {
	/* IS. terdefinisi array t yang berisi n harga barang yang dibeli
	FS. mengembalikan total harga barang yang terdapat pada array t */
	for i := 0; i < n; i++ {
		total = total + t[i]
	}
	return total
}

func promo(t belanja, n int, hp *float64) {
	/* IS. terdefinisi array t yang berisi n harga barang yang dibeli dan terurut
	ascending
	FS. hp berisi total harga setiap barang setelah dikurangi dengan diskonnya*/
	for i := 0; i < n; i++ {
		diskon := float64(t[i]) - float64(t[i])*float(i+1)/float(100)
		*hp = *hp + diskon
	}
}
