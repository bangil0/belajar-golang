package main

import (
  "fmt"
)

func main() {
  var data_siswa = map[string]string{}; // buat array assosiatif dengan tipe string
  data_siswa["nama"] = "Nama Siswa";
  data_siswa["kelas"] = "Kelas 1";
  data_siswa["umur"] = "10";
  
  fmt.Println(data_siswa);
  fmt.Printf("Nama : %v \n", data_siswa["nama"]);
  fmt.Printf("Kelas : %v \n", data_siswa["kelas"]);
  fmt.Printf("Umur : %v \n", data_siswa["umur"]);
}
