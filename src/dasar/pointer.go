package main

import (
	"fmt"
)

func test() {
	var nilai1 int = 10
	var penunjuk1 = &nilai1
	var penunjuk2 = &penunjuk1
	var penunjuk3 = &penunjuk2
	fmt.Println("=======================================")
	fmt.Println("Nilai 1 ", nilai1)
	fmt.Println("Lokasi variabel nilai 1 ", penunjuk1)
	fmt.Println("Lokasi variabel penunjuk 1 ", penunjuk2)
	fmt.Println("Lokasi variabel penunjuk 2 ", penunjuk3)
}


func main() {
	var umur int = 10
	var penunjuk = &umur // variabel penunjuk menyimpan alamat memori umur.
	fmt.Printf("Alamat umur %v \n", &umur)
	fmt.Printf("Umur sebelum diubah %v \n", umur)

	// ubah nilai umur lewat pointer
	*penunjuk = 21
	fmt.Printf("Umur sesudah diubah %v ", umur)

	test()
=

}
