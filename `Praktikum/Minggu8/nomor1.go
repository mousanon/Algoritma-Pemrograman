package main
import "fmt"

const N int = 127

func main() {

	var tabel[N]string
	var n int

	isiarray(&tabel, &n)
	fmt.Println("\nPalindrom? ", palindrom(tabel, n))
	
}

func isiarray(tabel *[N]string, n *int) {
	var x string
	var i int = 0 
	fmt.Scan(&x)
	for x != "." {
		tabel[i] = x
		fmt.Scan(&x)
		i++
		*n++
	}
}

func cetak(tabel [N]string, n int) {
	for i := 0; i < n; i++ {
		fmt.Print(tabel[i], " ")
	}
}


func balikan(tabel *[N]string, n int) {
	var temp string
	var awal int = 0
	var akhir int = n-1
	
	for awal != (n/2) {
		temp = tabel[awal]

		tabel[awal] = tabel[akhir]
		tabel[akhir] = temp
		awal++
		akhir--
	}
}


func palindrom(tabel [N]string, n int) bool {
	var balik[N]string 
	balik = tabel
	balikan(&balik, n)

	for i := 0; i < n-1; i++ {
		
		if tabel[i] != balik[i] {
			return false
		}
	}

	return true
}