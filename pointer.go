package main

import "fmt"

func main() {
  var umur int = 10;
  var penunjuk = &umur; // variabel penunjuk menyimpan alamat memori umur.
  fmt.Printf("Alamat umur %v \n", &umur);
  fmt.Printf("Umur sebelum diubah %v \n", umur);
  
  // ubah nilai umur lewat pointer
  *penunjuk = 21;
  fmt.Printf("Umur sesudah diubah %v ", umur);
}
