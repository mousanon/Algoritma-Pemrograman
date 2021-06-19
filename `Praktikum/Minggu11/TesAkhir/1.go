package main

import "fmt"

const NMAX = 100

type provinsi struct {
	nama string
	populasi int
	tumbuh float64
}

type dataProvinsi struct {
	type arrayProvinsi [NMAX]provinsi
	nProvinsi int
}

func main() {
	var arr arrayProvinsi
	var prov string
	var i int = 0

	fmt.Scanln(&arr[i].nama, &arr[i].populasi, &arr[i].tumbuh)
	for arr[i].nama != "NONE" && arr[i].populasi != 0 && arr[i].tumbuh != 0.0 {
		fmt.Scanln(&arr[i].nama, &arr[i].populasi, &arr[i].tumbuh)
		i++
	}

	fmt.Print("Nama provinsi? ")
	fmt.Scanln(prov)
	cari := cariProvinsi(arr, prov)
	fmt.Println(cari)

	fmt.Print("Prediksi populasi tahun depan provinsi? ")
	fmt.Scanln(prov)
	predict := prediksi(arr, prov)
	fmt.Println("Populasi Provinsi ", prov, "tahun depan: ", predict)
}

func cariProvinsi(data dataProvinsi, nama string) provinsi {
	for i := 0; i < NMAX; i++ {
		if nama == data[i].nama {
			return data[i].nama, data[i].populasi, data[i].tumbuh
		}
	}
}

func prediksi(data dataProvinsi, nama string) int {
	for i := 0; i < NMAX; i++ {
		if nama == data[i].nama {
			return data[i].populasi + data[i].populasi * data[i].tumbuh
		}
	}
}

func urutData(data *dataProvinsi) {
	for i := 0; i < NMAX; i++ {
		for j := i; j < n; j++ {
			var temp provinsi
			if arr[j].populasi > arr[i].populasi {
				temp = arr[j]
				arr[j] = arr[i]
				arr[i] = temp
			} 
		}

	}
	for i := 0; i < NMAX; i++ {
		fmt.Println(arr[i].nama, arr[i].populasi, arr[i].tumbuh)
	}
}