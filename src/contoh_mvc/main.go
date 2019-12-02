/*
 * JALANKAN PERINTAH GO BUILD DI FOLDER INI UNTUK MENGCOMPILE PROGRAM
 */

package main

import (
  "fmt"
  "contoh_mvc/controllers"
  "net/http"
  "strings"
  "flag"
)

func main() {  
  //~ var siswa models.Siswa;
  
  //~ // AMBIL DATA
  //~ var data_siswa []models.Siswa = siswa.GetData();
  //~ fmt.Println(data_siswa);
  
  //~ // ATUR DATA
  //~ siswa.SetNama("Nama dari models");
  //~ siswa.SetKelas("Kelas dari models");
  
  //~ // SIMPAN DATA
  //~ siswa.SaveData();
  
  //~ // ATUR DATA
  //~ siswa.SetId(10);
  //~ siswa.SetNama("Nama yang baru dari models");
  //~ siswa.SetKelas("Kelas baru dari models");
  
  //~ // UPDATE DATA
  //~ siswa.UpdateData();
  
  //~ // ATUR DATA
  //~ siswa.SetId(100);
  
  //~ // HAPUS DATA
  //~ siswa.DeleteData();
  
  //~ data_siswa = siswa.GetData();
  //~ fmt.Println(data_siswa);
  
  // ROUTE
  // route
  
  http.HandleFunc("/siswa", controllers.ListSiswa) 
  http.HandleFunc("/api/siswa", controllers.GetSiswa) 
  http.HandleFunc("/api/siswa/save", controllers.SaveSiswa) 
  
  directory := flag.String("d", ".", "the directory of static file to host")
	flag.Parse()
  http.Handle("/assets/", http.StripPrefix(strings.TrimRight("/assets/", "/"), http.FileServer(http.Dir(*directory))))
  fmt.Println("starting web server at http://localhost:8080/")
  http.ListenAndServe(":8080", nil)
  
}
