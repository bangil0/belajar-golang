package main

import (
  "fmt"
  "net/http" // package untuk web server
  _ "github.com/go-sql-driver/mysql" // import package db
  "encoding/json"
)

type Siswa struct {
  Id string `json:"id"`
  Nama string `json:"nama"`
  Kelas string `json:"kelas"`
}

// method untuk koneksi ke database
func db() (*sql.DB, error) {
    db, err := sql.Open("mysql", "root:mysql@tcp(db:3306)/mandanon_hpp") // terhubung ke database
    if err != nil { // cek error
        return nil, err
    }
    return db, nil // kembalikan instansi database
}

func ambilData(res http.ResponseWriter, req *http.Request) {
  db, err := db() // ambil koneksi database ke variabel
  if err != nil { // cek error
    fmt.Println(err.Error())
    return
  }
  defer db.Close() // tutup koneksi nanti 
  
  rows, err := db.Query("select * from siswa where ?", 1) // jalankan query sql
  if err != nil { // cek eksekusi error
      fmt.Println(err.Error())
      return
  }
  defer rows.Close() // tutup proses menjalankan querynya belakangan
  
  var data_siswa []Siswa // variabel penampung haasil database
  
  for rows.Next() { // for each
    var siswa = Siswa{} // buat variabel penampung satu baris
    var err = rows.Scan(&siswa.Id, &siswa.Nama, &siswa.Kelas) // ambil data setiap baris

    if err != nil { // cek error
        fmt.Println(err.Error())
        return
    }

    data_siswa = append(data_siswa, siswa) // masukkan hasil pembacaan ke array siswa
  }
  
  // buat variabel penampung json
  var json, _ = json.Marshal(data_siswa);
  
  // untuk menentukan header response
  res.Header().Set("Content-Type", "application/json")
  
  // hasilnya dikirim ke klien
  fmt.Fprintln(res, string(json));
  
}

func main() {
    // route
    http.HandleFunc("/api/siswa", ambilData) 

    fmt.Println("starting web server at http://localhost:8080/")
    http.ListenAndServe(":8080", nil)
}
