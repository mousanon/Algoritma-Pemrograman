package main
import "fmt"

func main() {
	var a, b, c, d, e byte
	var x int
	var f, g, h, i, j int

	fmt.Scanf("%c%c%c%c%c", &a, &b, &c, &d, &e)
	fmt.Scanf("%d", &x)

	f = int(a) + x
	g = int(b) + x
	h = int(c) + x
	i = int(d) + x
	j = int(e) + x

	fmt.Scanf("%c%c%c%c%c", f, g, h, i, j)
}