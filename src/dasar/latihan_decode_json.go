package main

import (
  "fmt"
  "encoding/json"
  "strings"
)


// struktur data jsonnya harus ditentukan terlebih dahulu
// Nama field di struct harus diawali huruf besar
type Mahasiswa struct {
  Nm_mhs string `json:"nm_mhs"`
  Tgl_lahir string `json:"tgl_lahir"`
  Jurusan string `json:"jurusan"`
}

func main() {
  // data json asli, berupa array of object
  var data_mhs_array = make([]string, 0);
  var data_mhs_json string
  var data_mhs []Mahasiswa
  
  // agar gampang, data mhsnya disusun jadi array dan nanti digabungkan dengan strings.join
  data_mhs_array = append(data_mhs_array, `{"nm_mhs": "Mahasiswa 1", "tgl_lahir": "2019-10-10", "jurusan": "SI"}`);
  data_mhs_array = append(data_mhs_array, `{"nm_mhs": "Mahasiswa 2", "tgl_lahir": "2019-11-10", "jurusan": "MI"}`);
  data_mhs_array = append(data_mhs_array, `{"nm_mhs": "Mahasiswa 3", "tgl_lahir": "2019-12-10", "jurusan": "TI"}`);
  
  // hasil string json dibuat
  data_mhs_json = fmt.Sprintf(`[%v]`, strings.Join(data_mhs_array, ","));
  
  
  // json.Unmarshal untuk decode json, variabel err untuk mengecek proses decode
  // hasil decode akan masuk ke variabel "data"
  var err = json.Unmarshal([]byte(data_mhs_json), &data_mhs)
  
  // cek proses error
  if err != nil {
    fmt.Println(err.Error())
    return
  }
  
  for _, mhs := range data_mhs {
    fmt.Printf("Nama: %v, Tanggal Lahir: %v, Jurusan: %v \n", mhs.Nm_mhs, mhs.Tgl_lahir, mhs.Jurusan);
  } 
  
}
