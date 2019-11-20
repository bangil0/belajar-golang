package main

import "fmt"

func main() {
  var daftar_siswa [2]string;
  var daftar_mahasiswa = []string{"Mhs 1", "Mhs 2", "Mhs 3", "Mhs 4"}
  daftar_siswa[0] = "Siswa 1";
  daftar_siswa[1] = "Siswa 2";
  fmt.Println(daftar_siswa[0]);
  fmt.Println(daftar_siswa[1]);
  fmt.Printf("Banyak siswa adalah %v \n", len(daftar_siswa));
  fmt.Println(daftar_mahasiswa[0]);
  fmt.Println(daftar_mahasiswa[1]);
  fmt.Println(daftar_mahasiswa[2]);
  fmt.Println(daftar_mahasiswa[3]);
  daftar_mahasiswa = append(daftar_mahasiswa, "Mhs baru"); // menambah isi array
  fmt.Println(daftar_mahasiswa[4]);
  
  // membaca isi array otomatis
  for no, mahasiswa := range daftar_mahasiswa {
    fmt.Printf("No %v, %v \n", no, mahasiswa);
  }
}
