package main
import "fmt"

func main() {
	var nilai int
	budget := 0
	fmt.Scanln(&nilai, &budget)
	if nilai > 80 {
		if nilai == 100 {
			budget += 100000
		} else if nilai > 85 && nilai < 99{
			budget += 50000
		} else if nilai <= 85 {
			budget += 25000
		}
	} else {
		budget += 0
	}
	fmt.Println(budget)
}