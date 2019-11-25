package main

import (
  "fmt"
  "os"
)

// argument disini maksudnya adalah parameter yang kita masukkan saat program dijalankan lewat command line
// contoh :
// ./nama_build_file satu dua tiga -> satu, dua dan tiga adalah argumentnya
// ./nama_build_file -nama=madam -kelas=satu -> parameternya berupa key value

func main() {
  var args = os.Args; // mengambil parameternya. hasilnya berupa array/slice
  fmt.Println(args);
}
