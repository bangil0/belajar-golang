package main

import "fmt"
import "net/http" // package untuk web server

func cekMethod(res http.ResponseWriter, req *http.Request) {
  
  // cara mengecek method web form
  // golang bisa mengenal semua method seperti put dan patch
  // req.FormValue("data") untuk mengambil data GET atau POST
  if req.Method  == "POST" {
    var data = req.FormValue("data"); // cara mengambil nilai FORM 
    fmt.Fprintln(res, "INI METHODNYA POST, datanya adalah ", data);
  } else if req.Method == "GET" {
    var data = req.FormValue("data"); // cara mengambil nilai FORM 
    fmt.Fprintln(res, "IN METHODNYA GET", data);
  } else {
    fmt.Fprintln(res, "Methodnya adalah ", string(req.Method))
  }
}

func main() {
    // route
    http.HandleFunc("/cek", cekMethod) 

    fmt.Println("starting web server at http://localhost:8080/")
    http.ListenAndServe(":8080", nil)
}
