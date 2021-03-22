package main

import "fmt"

func konversiSuhu(suhu *float64) {
	*suhu = (9/5 * (*suhu)) + 32
}

func main() {
	var suhu float64
	fmt.Scanln(&suhu)
	fmt.Println("Suhu dalam Celcius:", suhu)
	konversiSuhu(&suhu)
	fmt.Println("Suhu dalam Fahrenheit:", suhu)
}