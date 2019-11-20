package main

import "fmt"

func main() {
  var pilihan string = "Y";
  for pilihan == "Y" {
    fmt.Print("Jalankan aplikasi [Y/t]?");
    fmt.Scanln(&pilihan);
  }
  fmt.Print("Sampai jumpa! Tekan enter untuk keluar dari aplikasi");
  fmt.Scanln();
}
