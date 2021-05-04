package main
import "fmt"

const N int = 127

type tabGol struct {
	menang int
}

type ray = [N] tabGol

func main() {

	var pemain1, pemain2, pemain3 ray
	var a, b, c int

	input(&pemain1, &a)
	input(&pemain2, &b)
	input(&pemain3, &c)

	fmt.Print(rata(pemain1, a), rata(pemain2, b), rata(pemain3, c))
	
}

func input(t *ray, n *int) {

	var x int
	*n = 0
	fmt.Scan(&x)
	for x >= 0 {
		t[*n].menang = x
		*n++
		fmt.Scan(&x)
	}
}

func rata(t ray, n int) float64 {
	var rata float64
	for i := 0; i < n; i++ {
		rata = rata + float64(t[i].menang)
	}

	return rata/float64(n)
}
