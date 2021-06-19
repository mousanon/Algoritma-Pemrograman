package main

import "fmt"

type tag struct {
	topik  string
	banyak int
}
type tabTopic [100]string
type tabTag [100]tag

func mengecekKeberadaantopik(data1 tabTag, n1 int, topik string) (bool, int) {
	var i int

	for i = 0; i < n1; i++ {
		if data1[i].topik == topik {
			return true, i
		}
	}
	return false, i
}
func mengisiArrayTabTag(data1 *tabTag, n1 *int, data2 tabTopic, n2 int) {
	var cek bool
	var posisi int

	for i := 0; i < n2; i++ {
		cek, posisi = mengecekKeberadaantopik(*data1, *n1, data2[i])
		if cek == true {
			data1[posisi].banyak++
		} else {
			data1[posisi].topik = data2[i]
			data1[posisi].banyak++
			*n1++
		}
	}
}
func trendingTopic(data1 tabTag, n1 int) string {
	var max int
	max = 0
	for i := 0; i < n1; i++ {
		if data1[i].banyak > data1[max].banyak {
			max = i
		}
	}
	return data1[max].topik
}
func main() {
	var topik string
	var daftarTab tabTopic
	var daftarTag tabTag
	var n1 int
	var n2 int = 0
	//isi TabTopic
	for {
		_, err := fmt.Scan(&topik)
		if err != nil {
			break
		}
		daftarTab[n2] = topik
		n2++
	}

	mengisiArrayTabTag(&daftarTag, &n1, daftarTab, n2)
	fmt.Println(trendingTopic(daftarTag, n1))

}
