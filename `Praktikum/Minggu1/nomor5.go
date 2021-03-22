package main
import "fmt"

func main() {
    var str string
	fmt.Scanln(&str)
	for str != "selesai" {
		fmt.Println(str)
		fmt.Scanln(&str)
	}
}