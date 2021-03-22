package main
import "fmt"

func main(){
    var nilai int
    var rata2 float64
    n := 0
    jumlah :=0
    fmt.Scanln(&nilai)
    for nilai != -1{
        n ++
        jumlah += nilai
        fmt.Scanln(&nilai)
    }
    if n == 0 {
        rata2 = 0
    }else {
        rata2 = float64(float64(jumlah )/ float64(n))
    }
    fmt.Println(rata2)
}