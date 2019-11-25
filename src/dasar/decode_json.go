package main

import "fmt"
import "encoding/json"


// struktur data jsonnya harus ditentukan terlebih dahulu
// CATATAN:
// NAMA FIELD STRUCT HARUS DIAWALI HURUF BESAR
type User struct {
  FullName string `json:"Name"` // nama jsonnya harus ditentukan juga di struct
  Age      int
}

func main() {
  // data json asli
  var jsonString = `{"Name": "john wick", "Age": 27}`
  // data jsonnya harus diubah kedalam byte terlebih dahulu
  var jsonData = []byte(jsonString)
  
  // variabel penampung data hasil json
  var data User
  
  // json.Unmarshal untuk decode json
  var err = json.Unmarshal(jsonData, &data)
  
  // cek proses error
  if err != nil {
      fmt.Println(err.Error())
      return
  }
  
  // menampilkan hasilnya
  fmt.Println("user :", data.FullName)
  fmt.Println("age  :", data.Age)
}
