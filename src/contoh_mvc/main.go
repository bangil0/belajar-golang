/*
 * JALANKAN PERINTAH GO BUILD DI FOLDER INI UNTUK MENGCOMPILE PROGRAM
 */

package main

import (
  "fmt"
  "contoh_mvc/models"
)

func main() {
  var madam = models.Siswa{};
  madam.SetNama("Madam import dari amrik");
  madam.SetKelas("Kelas Madam");
  madam.SetUmur(21);
  fmt.Println(madam.GetNama());
  fmt.Println(madam.GetKelas());
  fmt.Println(madam.GetUmur());
}
