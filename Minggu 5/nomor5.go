package main

import "fmt"

func is_prime(angka int) int {
  count := 0
  if angka == 2 {
    count++
  } else if angka / angka == 1 && angka / 1 == angka {
    count++
  }
  return count
}

func main() {
  var bilangan, jumlahPrima int
  for bilangan != 0 {
    fmt.Scanln(&bilangan)
  }
  jumlahPrima = is_prime(bilangan)
  fmt.Print("Banyak prima : ", jumlahPrima)
}
