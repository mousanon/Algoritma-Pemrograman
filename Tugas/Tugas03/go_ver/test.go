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
