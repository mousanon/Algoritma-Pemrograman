package main

import "fmt"

const NMAX int = 10

type lagu struct {
	judul    string
	penyanyi string
	durasi   waktu
}

type waktu struct {
	menit int
	detik int
}

type playlist [NMAX]lagu

func searchSong(T playlist, n int, judul string, penyanyi string) bool {
	for i := 0; i < n; i++ {
		if judul == T[i].judul && penyanyi == T[i].penyanyi {
			return true
		}
	}
	return false
}
func buatPlaylist(T *playlist, n *int) {
	var (
		penyanyi, judul string
		menit, detik    int
		check           bool
	)

	fmt.Scanln(&judul, &penyanyi, &menit, &detik)
	for judul != "#" && penyanyi != "#" {
		check = searchSong(*T, *n, judul, penyanyi)
		if check == false {
			T[*n].judul = judul
			T[*n].penyanyi = penyanyi
			T[*n].durasi.menit = menit
			T[*n].durasi.detik = detik
			*n++
		}
		fmt.Scanln(&judul, &penyanyi, &menit, &detik)
	}
}
func durasiLama(T playlist, n int) int {
	var maxMenit int
	var maxDetik int
	var maxDurasi int

	maxMenit = 0
	maxDetik = 0
	for i := 0; i < n; i++ {
		fmt.Println(maxMenit, T[i].durasi.menit)
		if maxMenit < T[i].durasi.menit {
			maxDurasi = i
			maxMenit = T[i].durasi.menit
			maxDetik = T[i].durasi.detik
		} else if maxMenit == T[i].durasi.menit {
			if maxDetik < T[i].durasi.detik {
				maxDurasi = i
				maxMenit = T[i].durasi.menit
				maxDetik = T[i].durasi.detik
			}
		}
	}
	return maxDurasi
}
func cetakPlaylist(T playlist, n int) {
	var laguTerlama int
	laguTerlama = durasiLama(T, n)
	for i := 0; i < n; i++ {
		if i == laguTerlama {
			fmt.Println("*"+T[i].judul, T[i].penyanyi, T[i].durasi.menit, T[i].durasi.detik)
		} else {
			fmt.Println(T[i].judul)
		}
	}
}
func main() {
	var data playlist
	var n int

	buatPlaylist(&data, &n)
	cetakPlaylist(data, n)
}
