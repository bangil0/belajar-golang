package main

import (
  "time"
  "fmt"
)

// anggap saja fungsi ini prosesnya mirip mengambil data dari database
func ambilData(banyak_pesan int) (hasil []string) {
  for x := 0; x < banyak_pesan; x++ {
    waktu := time.Now();
    hasil = append(hasil, fmt.Sprintf("[%v:%v:%v] go routine. %v \n ", waktu.Hour(), waktu.Minute(), waktu.Second() , x)); // kumpulkan pesan terlebih dahuli kevariabel lain
  }
  return;
}

func main() {
  var data_pertama = make(chan []string, 0); // variabel ini dinamakan chanel. berguna untuk menampung data dari proses go routine atau proses yang sync
  var data_kedua = ambilData(100000);
  go func () {
    data_pertama <- ambilData(100000); // memindahkan isi proses async ke chanel, menggunakan tanda <- disebelah kanan
  }();
  fmt.Print(<- data_pertama); // sedangkan mengambil isi channel tanda <- diposisikan disebelah kiri
  fmt.Print(data_kedua);
}
