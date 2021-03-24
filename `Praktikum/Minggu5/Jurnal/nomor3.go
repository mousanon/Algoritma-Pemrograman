package main

import {
    "fmt"
    "math"
}


func jarak(a,b,c,d float64) float64 {    // function jarak
    return math.sqrt((a - c) * (a - c) + (b - d) * (b - d))    // me-return perhitungan rumus jarak 
}

func posisi(cx,cy,r,x,y float64) bool {    // function posisi
    var jarak1 float64 = jarak(cx,cy,x,y)    // mendeklarasikan variabel lokal jarak1 dengan tipe data float64 dan memanggil function jarak

    return jarak1 < r    // me-return apakah jarak kurang dari r
}

func main() {    // program utama
    var cx,cy,x,y,r float64   // mendeklarasikan variabel cx. cy, x, y, r dengan tipe data float 64
    var dalam bool    // mendeklarasikan variabel dalam dengan tipe data boolean

    fmt.Scanln(&cx,&cy,&r)    // mrnginpuyksn cx, cy dan r
    fmt.Scanln(&x,&y)    // menginputkan x dan y

    dalam = posisi(cx,cy,r,x,y)    // memanggil function posisi dan dimasukkan variabel dalam

    if dalam {    // jika variabel dalam memenuhi lakukan program dibawah
        fmt.Println("Anda berada di dalam taman")    // menampikan ke layar jika if statement benar
    } else {    // jika if statement tidak terpenuhi, maka menjalankan program dibawah
        fmt.Println("Anda berada di luar taman")    // menampilkan ke layar jika else memenuhi
    }

}