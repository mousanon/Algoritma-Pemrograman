package main
import "fmt"

func main() {
	var benar, tb1, tb2, tb3 int
	var a, b, c, d int
	var x, y, z int

	fmt.Scanln(&benar)
	fmt.Scanln(&tb1)
	fmt.Scanln(&tb2)
	fmt.Scanln(&tb3)

	a = benar / 1000
	b = benar / 100 % 10
	c = benar % 100 / 10
	d = benar % 10

	x = benar / 100
	y = benar % 1000 / 10
	z = benar % 100

	fmt.Println((tb1 == a) || (tb1 == b) || (tb1 == c) || (tb1 == d))
	fmt.Println((tb2 == x) || (tb2 == y) || (tb2 == z))
	fmt.Println(tb3 == benar)

}