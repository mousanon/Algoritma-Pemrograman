package main
import "fmt"

func main() {
	var n, a, b, c int   // mendeklarasikan variabel n, a, b dan c dalam bentuk integer
	fmt.Scanln(&n)   // meminta input n berupa range dari peruangan

	i := 0   // mendeklarasikan nilai i adalah 0
	for i < n {   // selama i di dalam range n lakukan program dibawah
		fmt.Scanln(&a, &b)   // meminta input untuk variabel a dan b
		c = a + b    // emnjumlahkan nilai a dan b dan menyimpannya di variabel c
		if c % 2 == 0 {   // jika modulo dari c dari 2 sama dengan 0 lakukan program dibawah
			fmt.Println(a)    // menampilkan nilai a ke layar
		} else {   // jika if diatas tidak terpenuhi akan masuk ke else dan menjalankan program dibawah
			fmt.Println(b)    // menampilkan nilai b ke layar
		}
		i++   // menambhkan i 1kali setiap selesai dengan for loop
	}
}