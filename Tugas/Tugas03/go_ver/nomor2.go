package main

import "fmt"

type tabStr [100]string

func main() {
	var t tabStr
	var teks, w string
	var k, posisi, idx int

	fmt.Scanln(&teks, &w)
	isiArray(teks, &t, &k)
	cariPosisi(t, k, w, &posisi, &idx)
}

func panjang(s string) int {
	return len(s)
}

func isiArray(s string, daftar *tabStr, k int) {
	var j int
	j = panjang(s)
	for i := 0; i < j; i++ {
		if string(s[i]) != "_" {
			daftar[k] += string(s[i])
		} else {
			*k++
		}
	}
}

func indexKata(kata, w string) int {
	var idx int
	for i := 0; i < panjang(kata); i++ {
		if string(w[0]) == string(kata[i]) {
			idx = i
		} else {
			idx = -1
		}
	}
	return idx
}

func cariPosisi(daftar tabStr, n int, w string, posisi, idx *int) {
	var i int
	*idx = -1
	for i <= n {
		for j := 0; j < panjang(daftar[i]); j++ {
			if indexKata(string(daftar[i][j]), w) != -1 {
				*idx = j
				*posisi = i + 1
			}
		}
		i++
	}
	if idx != -1 {
		fmt.Println(posisi, *idx)
	} else {
		fmt.Println(-1)
	}
}
