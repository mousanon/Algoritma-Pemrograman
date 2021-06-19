package main

import "fmt"

type arrLagu = []string
type arrPenyanyi = []string
type arrMenit = []int
type arrDetik = []int

func main() {
	var lagu arrLagu
	var penyanyi arrPenyanyi
	var menit arrMenit
	var detik arrDetik
	var song, singer string
	var minutes, seconds, valid int
	var i int = 0

	fmt.Scan(&song, &singer)
	fmt.Scanln(&minutes, &seconds)
	for song != "" && singer != "" {
		for song != "#" && singer != "#" {
			valid++
			lagu[i] = song
			penyanyi[i] = singer
			menit[i] = minutes
			detik[i] = seconds
			i++
		}
		fmt.Scan(&song, &singer)
		fmt.Scanln(&minutes, &seconds)
	}
	searching(lagu, penyanyi, song, singer, valid)
}

func searching(lagu arrLagu, penyanyi arrPenyanyi, song, singer string, valid int) {
	for i := 0; i < valid; i++ {
		if lagu[i] == song {
			break
		}
	}
	for j := 0; j < valid; j++ {
		if penyanyi[j] == singer {
			break
		}
	}
	fmt.Println(song)
}
