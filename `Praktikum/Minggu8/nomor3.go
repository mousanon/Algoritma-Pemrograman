package main
import "fmt"

const nMax = 10

type Buku struct{
	judul, penulis string
	tahun int
}

type TabBuku = [nMax]Buku

func tambahBuku(kardus *TabBuku, atas *int){
	var (
		judul,penulis string
		tahun,n int
	)
	fmt.Print("\nTambah berapa buku: ")
	fmt.Scanln(&n)
	i:=0
	for i<n && *atas<=nMax-1{
		*atas++
		fmt.Print("Judul Buku ke-", i+1, ": ")	
		fmt.Scanln(&judul)
		fmt.Print("Penulis: ")
		fmt.Scanln(&penulis)
		fmt.Print("Tahun: ")
		fmt.Scanln(&tahun)
		kardus[i].judul = judul
		kardus[i].penulis = penulis
		kardus[i].tahun = tahun
		i++	
	}
}

func ambilBuku(kardus *TabBuku, atas *int, ambil *Buku){
	*ambil = kardus[*atas]
	*atas--
}

func cariBuku(kardus *TabBuku, atas *int, X string){
	var (
		found bool
		ambil Buku
	)
	found = false
	i:=*atas
	
	for i>=0 && !found {
		ambilBuku(kardus,&i,&ambil)

		if (ambil.judul!=X){
			fmt.Print(ambil.judul," ")
		} else {
			fmt.Println("KETEMU")
			found = true
		}
	}
	if found == false {
		fmt.Printf("\nBUKU DENGAN JUDUL %s TIDAK KETEMU",X)
	}
}

func main(){
	var (
		kardus TabBuku
		judul string
		atas int
	)
	// kardus kosong
	atas = -1

	// Tambah 4 Buku, n = 4
	tambahBuku(&kardus,&atas)

	// cari buku dengan judul = C
	fmt.Print("Cari buku dengan Judul: ")
	fmt.Scanln(&judul)
	cariBuku(&kardus,&atas,judul)

	// Tambah buku sampai penuh, 1 lagi n = 1
	tambahBuku(&kardus,&atas)

	// Cari buku dengan judul yang tidak terdapat pada kardus tersebut
	fmt.Print("Cari buku dengan Judul: ")
	fmt.Scanln(&judul)
	cariBuku(&kardus,&atas,judul)

}