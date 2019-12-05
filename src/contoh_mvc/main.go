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
  // ROUTE
  // route
  
  http.HandleFunc("/siswa", controllers.ListSiswa) 
  http.HandleFunc("/siswa/edit", controllers.EditSiswa) 
  http.HandleFunc("/siswa/add", controllers.AddSiswa) 
  http.HandleFunc("/siswa/delete", controllers.DeleteSiswa) 
  
  // ROUTE UNTUK FILE STATIS
  http.Handle("/assets/",
        http.StripPrefix("/assets/",
            http.FileServer(http.Dir("assets"))))
            
  fmt.Println("starting web server at http://localhost:8080/")
  http.ListenAndServe(":8080", nil)
  
}
