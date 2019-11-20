package main

import "fmt"

// buat dulu struktur barisnya
type Siswa struct {
  nama, kelas string
  umur int
}

func main() {
  var data_siswa = []Siswa{
    Siswa{"Nama Siswa 1", "XII1", 24},
    Siswa{"Nama Siswa 2", "XII2", 23},
    Siswa{"Nama Siswa 3", "XII3", 22},
  };
  
  var allStudents = []struct { // membuat array of object tanpa harus bikin struct diluar fungsi main
    grade int
  }{
    {grade: 2},
    {grade: 3},
    {grade: 3},
  }
  // menampilkan keseluruhan jadi string
  fmt.Println(data_siswa);
  
  // menampilkan pakai for
  fmt.Println("========================");
  fmt.Println("       DATA SISWA");
  fmt.Println("========================");
  fmt.Println("No | Nama Siswa      | Kelas   | Umur");
  for no, data := range data_siswa { // mirip foreach di PHP
    fmt.Printf("%v | %v      | %v   | %v \n", (no+1), data.nama, data.kelas, data.umur);
  }
  
}
