package main

import "fmt"

type tag struct {
	topik  string
	banyak int
}

func main() {
	var tabTopic [100]string
	var tabTag [100]tag
	var trendingtopic string

	for i := 0; i < 100; i++ {
		fmt.Scan(&tabTopic[i])
		if tabTopic[i] == "." {
			tabTopic[i] = ""
			break
		}
	}
	isi(tabTopic, &tabTag)
	topic(tabTag, &trendingtopic)

	fmt.Println(trendingtopic)

}

func cek(tabTopic string, tabTag [100]tag) int {
	var hasil int
	hasil = -1
	if tabTopic != "" {
		for i := 0; i < 100; i++ {
			if tabTopic == tabTag[i].topik {
				hasil = i
				break
			} else {
				hasil = -1
			}
		}
	}
	return hasil
}

func isi(tabTopic [100]string, tabTag *[100]tag) {
	var isi int
	isi = 0

	for i := 0; i < 100; i++ {
		if cek(tabTopic[i], *tabTag) >= 0 {
			tabTag[cek(tabTopic[i], *tabTag)].banyak++

		} else if cek(tabTopic[i], *tabTag) == -1 {
			tabTag[isi].topik = tabTopic[i]
			tabTag[isi].banyak++
			isi++
		}
	}
}

func topic(tabTag [100]tag, trendingtopic *string) {
	var topikterbanyak int
	topikterbanyak = -1
	for i := 0; i < 100; i++ {
		if topikterbanyak < tabTag[i].banyak {
			topikterbanyak = tabTag[i].banyak
			*trendingtopic = tabTag[i].topik
		}
	}
}
