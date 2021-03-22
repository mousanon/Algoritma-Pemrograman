package main

import "fmt"

func main() {
	var r, t, hitungVolume float64

	pi := 3.14

	fmt.Scanln(&r, &t)
 	hitungVolume = ((r * r) * pi) * t

	fmt.Println(hitungVolume)

}