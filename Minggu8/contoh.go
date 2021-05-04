package main

import "fmt"

type balok struct {
	p, l, t int 
	luas, volume int 
}

func main() {
	var cube balok
	fmt.Scanln(&cube.p, &cube.l, &cube.t)
	cube.volume = cube.p * cube.l * cube.t
	cube.luas = 2 * ((cube.p * cube.l) + (cube.p * cube.t) + (cube.l * cube.t))
	fmt.Println(cube.p, cube.l, cube.t, cube.luas, cube.volume)
}