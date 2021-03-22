package main

import "fmt"

func main() {
	var tendangan int  // meneta

	count := 1
	gol := 0
	// fmt.Print("Tendangan ke-", count, " : ")
	// fmt.Scanln(&tendangan)
	for tendangan > 0 {
		fmt.Print("Tendangan ke-", count, " : ")
		fmt.Scanln(&tendangan)
		if tendangan % 4 == 0 {
			gol++
		} 
		count++
	}
	fmt.Print("Skor gol: ", gol, " dari ", count, "tendangan")

}