package main
import "fmt"

func main() {
	var r, luas float64
	fmt.Scanln(&r, &luas)
	luas = (22 * r) / 7
	fmt.Println("Luas lingkaran dengan jari-jari =", r, "adalah", luas)
}