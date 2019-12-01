/*
 * JALANKAN PERINTAH GO BUILD DI FOLDER INI UNTUK MENGCOMPILE PROGRAM
 */

package main

import (
  "fmt"
  "contoh_mvc/models"
)

func main() {  
  var siswa models.Siswa;
  
  // AMBIL DATA
  var data_siswa []models.Siswa = siswa.GetData();
  fmt.Println(data_siswa);
  
  // ATUR DATA
  siswa.SetNama("Nama dari models");
  siswa.SetKelas("Kelas dari models");
  
  // SIMPAN DATA
  siswa.SaveData();
  
  // ATUR DATA
  siswa.SetId(10);
  siswa.SetNama("Nama yang baru dari models");
  siswa.SetKelas("Kelas baru dari models");
  
  // UPDATE DATA
  siswa.UpdateData();
  
  // ATUR DATA
  siswa.SetId(100);
  
  // HAPUS DATA
  siswa.DeleteData();
  
  data_siswa = siswa.GetData();
  fmt.Println(data_siswa);
  
}
