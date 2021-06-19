package main

import "fmt"

const NMAX = 10000

type lagu struct {
	judul, penyanyi string
	menit, detik    int
}

type playlist [NMAX]lagu

func main() {
	var T playlist
	var n int

	buatPlaylist(&T,&n)
	cetak(T, n)
}

func buatPlaylist(T *playlist, n *int) {
	var judul, penyanyi string
	var menit, detik int

	fmt.Scan(&judul, &penyanyi, &menit, &detik)
	for judul != "#" && penyanyi != "#" {
		if !cariLagu(*T, *n, judul, penyanyi) {
			T[*n].judul = judul
			T[*n].penyanyi = penyanyi
			T[*n].menit = menit
			T[*n].detik = detik
			*n++
		}
		fmt.Scanln(&judul, &penyanyi, &menit, &detik)
	}
}

func cariLagu(T playlist, n int, judul, penyanyi string) bool {
	for i := 0; i < n; i++ {
		if T[i].judul == judul && T[i].penyanyi == penyanyi {
			return true
		}
	}
	return false
}

func terlama(T playlist, n int) int {
	var lama int = 0
	for i := 0; i < n; i++ {
		if (T[i].menit * 60 + T[i].detik) > (T[lama].menit + T[lama].detik) {
			lama = i
		}
	}
	return lama
}

func cetak(T playlist, n int) {
	var palingLama int
	palingLama = terlama(T, n)
	for i := 0; i < n; i++ {
		if i == palingLama {
			fmt.Println("*", T[i].judul, ..)
		} else {
			fmt.Println(T.[i].judul)
		}
	}
}