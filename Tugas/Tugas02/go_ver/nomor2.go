package main

import (
	"fmt"
)

func main() {
	var x int64
	fmt.Scanln(&x)
	fmt.Println(des2bin(x))

}

func des2bin(num int64) (str string) {
	str = ""
	for !(num == 0) {
		if num%2 == 0 {
			str += "0"
			num = num / 2
		} else {
			str += "1"
			num = num / 2
		}
	}
	return reverse(str)
}
func reverse(str string) (result string) {
	for _, v := range str {
		result = string(v) + result
	}
	return
}