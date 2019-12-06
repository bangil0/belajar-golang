/*
 * JALANKAN PERINTAH GO BUILD DI FOLDER INI UNTUK MENGCOMPILE PROGRAM
 */

package main

import (
  "fmt"
  "contoh_mvc/controllers"
  "net/http"
)

func main() {
  var SiswaController controllers.SiswaController = controllers.NewSiswaController();
  // ROUTE
  // route
  
  http.HandleFunc("/siswa", SiswaController.ListSiswa) 
  http.HandleFunc("/siswa/edit", SiswaController.EditSiswa) 
  http.HandleFunc("/siswa/add", SiswaController.AddSiswa) 
  http.HandleFunc("/siswa/delete", SiswaController.DeleteSiswa) 
  
  // ROUTE UNTUK FILE STATIS
  http.Handle("/assets/",
        http.StripPrefix("/assets/",
            http.FileServer(http.Dir("assets"))))
            
  fmt.Println("starting web server at http://localhost:8080/")
  http.ListenAndServe(":8080", nil)
  
}
