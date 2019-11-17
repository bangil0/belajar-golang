package main

import "fmt"
type Siswa struct {
  nama, kelas string
  umur int
}

func main() {
  var madam = Siswa{"Nama", "Kelas", 12};
  fmt.Println(madam.nama);
  fmt.Println(madam.kelas);
  fmt.Println(madam.umur);
}
