package main

import "fmt"
import "net/http" // package untuk web server

func cekMethod(res http.ResponseWriter, req *http.Request) {
  
  // cara mengecek method web form
  // golang bisa mengenal semua method seperti put dan patch
  // req.FormValue("data") untuk mengambil data GET atau POST
  // req.FormValue MENGEMBALIKAN NILAI JADI STRING
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

func cekTipeData(res http.ResponseWriter, req *http.Request) {
  var data = req.FormValue("data");
  if data == "" {
    fmt.Fprintln(res, "Nilainya kosong");
    return
  }
  var hasil = fmt.Sprintf("Tipe data %T , nilainya %v", data, data);
  fmt.Fprintln(res, hasil);
}

func main() {
    // route
    http.HandleFunc("/cek", cekMethod) 
    http.HandleFunc("/cektipedata", cekTipeData) 

    fmt.Println("starting web server at http://localhost:8080/")
    http.ListenAndServe(":8080", nil)
}
