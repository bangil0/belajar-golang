package main

import (
  "fmt"
  "os/exec"
)

func main() {
  var pilihan int;
  var data = make([]string, 0);
  
  for pilihan != 3 {
    exec.Command("clear").Output();
    fmt.Println("=================================");
    fmt.Println("==============  MENU  ===========");
    fmt.Println("=================================");
    fmt.Println("1. Lihat Data");
    fmt.Println("2. Tambah Data");
    fmt.Println("3. Keluar");
    fmt.Println("Silahkan tentukan pilihan [1/2/3] ? ");
    fmt.Scanln(&pilihan);
    if pilihan == 2 {
      exec.Command("clear").Output();
      var data_tmp string;
      fmt.Println("Ketik data baru: ");
      fmt.Scanln(&data_tmp);
      data = append(data, data_tmp);
      fmt.Println("Data berhasil ditambahkan");
      fmt.Scanln();
    } else if pilihan == 1 {
      exec.Command("clear").Output();
      fmt.Println("Daftar Data : ");
      for no, dt := range data {
        fmt.Printf("%v. %v \n", no, dt);
      }
      fmt.Println("Tekan enter untuk kembali ke menu");
      fmt.Scanln();
    }
  }
  fmt.Println("Sampai Jumpa!");
}
