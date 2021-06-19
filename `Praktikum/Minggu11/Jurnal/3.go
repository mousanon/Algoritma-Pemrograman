package main 

import "fmt"

type tabdate[100000]string

func main() {
	var t tabdate
	var n int
}

func add(t *tabdate, n *int) {
	var inputan string
	fmt.Scan(&inputan)
	for inputan != "####.##.##" {
		t[*n] = inputan
		*n++
		fmt.Scan(&inputan)
	}

}

func urut(t *tabdate, n int) {
	var max int
	for i := 0; i < n; i++ {
		
	}
}

func maxpos()