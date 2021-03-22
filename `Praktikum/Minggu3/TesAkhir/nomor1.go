package main
import "fmt"

func main() {
	var n, i int    // mendeklarasikan variabel n dan i berupa integer
	var count int   // mendeklarasikan variabel count berupa integer
	fmt.Scanln(&n)   // meminta input untuk variabel n

	i := 1   // mendeklarasikan nilai i adalah 1
	count := 0    // mendeklarasikan nilai i adalah 0
	for i <= n {   // selama i lebih kecil atau sama dengan n lakukan program dibawah
		if n % i == 0{   // jika n dimodulo dengan i sama dengan 0 lakukan program diwabah
			fmt.Print(i, " ")    // menampilkan i ke layar
			count = count + 1    // menambah count 1
		}
		i++   // menambahkan iterasi
	}
	if count > 2 {   // jika count lebih dari 2 lakukan program dibawah
		fmt.Println("\nBukan Oleole")    // menampilkan ke layar
	} else {    // jika program if tidak memenuhi akan dijalankan dan menjalankan program dibawah
		fmt.Println("\nOleole")   // tampilkan ke  layar
		}

}