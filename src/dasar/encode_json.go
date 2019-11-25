package main

import "fmt"
import "encoding/json"

func main() {
  // ini adalah data yang akan diubah jadi json memakai struct
  var data_mahasiswa = []struct{
    Nama string `json:"nm_mhs"`
    Kelas string `json:"kelas"`
    Umur int `json:"umur"`
  }{
    {Nama: "Mhs 1", Kelas: "Kelas 1", Umur: 123},
    {Nama: "Mhs 2", Kelas: "Kelas 2", Umur: 124},
    {Nama: "Mhs 3", Kelas: "Kelas 3", Umur: 125},
  }
  
  // data dengan tipe map dan interface
  var data_siswa = []map[string]interface{} {
    {"nama": "Siswa 1", "kelas": "A1", "umur": 12},
    {"nama": "Siswa 2", "kelas": "A1", "umur": 11},
    {"nama": "Siswa 3", "kelas": "A1", "umur": 13},
  }
  
  // proses encode
  var mhs_json, err = json.Marshal(data_mahasiswa);
  var json_siswa, _ = json.Marshal(data_siswa);
   
  // cek error
  if err != nil {
    panic(err);
    return
  }
  
  // tampilkan hasil jsonnya
  fmt.Println(string(mhs_json));
  fmt.Println(string(json_siswa));
}
