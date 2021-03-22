package main
import "fmt"

func main() {
	var n int
	var index string
	var nilai, sks, jumlah int

	if index == "A" {
		nilai = 4
	} else if index == "B" {
		nilai = 3
	} else if index == "C" {
		nilai = 2
	} else if index == "D" {
		nilai = 1
	} else if index == "E" {
		nilai = 0
	}
	
	fmt.Scanln(&n)
	for n := 0; n <= n; n++ {
		fmt.Scanln(&index, &sks)
		jumlah = nilai*sks
	}
	fmt.Println(jumlah)

}