package main
import "fmt"

func main() {
	var secret, a, b int   // mendeklarasikan variabel secret, a dan b dalam bentuk integer
	fmt.Scanln(&secret)   // menginputkan secret

	for a != -1 && b != 1 {   // melakukan perulangan selama a tidak sama dengan -1 dan 1
		fmt.Scanln(&a, &b)   // menginputkan integer a dan b
		if a%secret == 0 {   // jika d di modulo dengan secret sama dengan 0
			fmt.Println(b)   // tampilkan b di layar
		} else if b % secret == 0 {   // jika if statement diatas tidak memenuhi akan masuk ke else if yaitu jika b di modulo dengan secret sama dengan 0
			fmt.Println(a)   // tampilkan b di layar
		}
	}
}