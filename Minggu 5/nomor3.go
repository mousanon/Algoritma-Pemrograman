package main

import "fmt"

func diskon(t_blnja int, member bool) float64 {
	member = false
	if t_blnja > 100000 && !member {
		return t_blnja * 0.1
	} else if t_blnja > 100000 && member  {
		return t_blnja * 0.05
	}  else {
		return 0
	}
}

func main() {
	var t_blnja, total int
	var member bool
	fmt.Scanln(&t_blnja, &member)
	total = t_blnja - diskon(t_blnja, member)
	fmt.Print("Total bayar : ", total)
}
