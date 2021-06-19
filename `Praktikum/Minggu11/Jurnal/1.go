package main

import "fmt"

type bola struct {
	nama1, nama2 string
	gol, assist  int
}

type tabGoal [100]bola

func main() {
	var arr tabGoal
	var n int

	fmt.Scanln(&n)
	for i := 0; i < n; i++ {
		fmt.Scan(&arr[i].nama1, &arr[i].nama2, &arr[i].gol, &arr[i].assist)
	}

	for i := 0; i < n; i++ {
		for j := i; j < n; j++ {
			var temp bola
			if arr[j].gol > arr[i].gol {
				temp = arr[j]
				arr[j] = arr[i]
				arr[i] = temp
			} else if arr[j].gol > arr[i].gol {
				if arr[j].assist > arr[i].assist {
					temp = arr[j]
					arr[j] = arr[i]
					arr[i] = temp
				}
			}
		}

	}
	for i := 0; i < n; i++ {
		fmt.Println(arr[i].nama1, "", arr[i].nama2, "", arr[i].gol, "", arr[i].assist)
	}
}
