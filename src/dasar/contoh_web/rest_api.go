package main

import "fmt"
import "net/http" // package untuk web server
import "encoding/json"

type Siswa struct {
  Nm_siswa string `json:"nm_siswa"`
  Kelas string `json:"kelas"`
  Umur int `json:"umur"`
}

func getSiswa(res http.ResponseWriter, req *http.Request) {
  // untuk menentukan header response
  res.Header().Set("Content-Type", "application/json")
  
  var data_siswa = []Siswa{
    {Nm_siswa: "Siswa 1", Kelas: "XII-IPA 4", Umur: 18},
    {Nm_siswa: "Siswa 11", Kelas: "XII-IPA 2", Umur: 19},
    {Nm_siswa: "Siswa 311", Kelas: "XII-IPA 1", Umur: 20},
    {Nm_siswa: "Siswa 21", Kelas: "XII-IPA 4", Umur: 21},
    {Nm_siswa: "Siswa 14", Kelas: "XII-IPA 3", Umur: 22},
    {Nm_siswa: "Siswa 51", Kelas: "XII-IPA 5", Umur: 23},
    {Nm_siswa: "Siswa 112", Kelas: "XII-IPA 7", Umur: 24},
  }
  
  var data_web, _ = json.Marshal(data_siswa);
  fmt.Fprintln(res, string(data_web));
}

func main() {
    // route
    http.HandleFunc("/api/siswa", getSiswa) 

    fmt.Println("starting web server at http://localhost:8080/")
    http.ListenAndServe(":8080", nil)
}
