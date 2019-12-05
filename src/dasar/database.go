package main

import (
  "fmt"
  "database/sql" // datase
  _ "github.com/go-sql-driver/mysql" // import package db
  "encoding/json"
)

type Siswa struct {
  Id int `json:"id"`
  Nama string `json:"nama"`
  Kelas string `json:"kelas"`
}

// method untuk koneksi ke database
func db() (*sql.DB, error) {
    db, err := sql.Open("mysql", "root:mysql@tcp(localhost:3306)/golang") // terhubung ke database
    if err != nil { // cek error
        return nil, err
    }
    return db, nil // kembalikan instansi database
}


func getData() (data_siswa []Siswa) {
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
  
  for rows.Next() { // for each
    var siswa = Siswa{} // buat variabel penampung satu baris
    var err = rows.Scan(&siswa.Id, &siswa.Nama, &siswa.Kelas) // ambil data setiap baris

    if err != nil { // cek error
        fmt.Println(err.Error())
        return
    }

    data_siswa = append(data_siswa, siswa) // masukkan hasil pembacaan ke array siswa
  }
  return
}

func saveData(data_siswa Siswa) {
  db, err := db() // ambil koneksi database ke variabel
  if err != nil { // cek error
    fmt.Println(err.Error())
    return
  }
  defer db.Close() // tutup koneksi nanti 
  
  hasil, err := db.Exec("INSERT INTO siswa (nama, kelas) VALUES(?, ?)", data_siswa.Nama, data_siswa.Kelas);
  if err != nil {
    fmt.Println(err);
    return
  }
  last_id, _ := hasil.LastInsertId();  // AMBI LAST INSERT ID
  rows_affected, _ := hasil.RowsAffected(); // AMBIL BARIS YANG TERPENGARUH
  
  fmt.Println("LAST INSERT ID ", last_id);
  fmt.Println("AFFECTED ROWS ", rows_affected);
}

func updateData(data_siswa Siswa) {
  db, err := db() // ambil koneksi database ke variabel
  if err != nil { // cek error
    fmt.Println(err.Error())
    return
  }
  defer db.Close() // tutup koneksi nanti 
  
  hasil, err := db.Exec("UPDATE siswa SET nama = ?, kelas = ? WHERE id = ?", data_siswa.Nama, data_siswa.Kelas, data_siswa.Id);
  if err != nil {
    fmt.Println(err);
    return
  }
  rows_affected, _ := hasil.RowsAffected(); // AMBIL BARIS YANG TERPENGARUH

  fmt.Println("AFFECTED ROWS ", rows_affected);
}
func deleteData(id string) {
  db, err := db() // ambil koneksi database ke variabel
  if err != nil { // cek error
    fmt.Println(err.Error())
    return
  }
  defer db.Close() // tutup koneksi nanti 
  
  hasil, err := db.Exec("DELETE FROM siswa WHERE id = ?", id);
  if err != nil {
    fmt.Println(err);
    return
  }
  rows_affected, _ := hasil.RowsAffected(); // AMBIL BARIS YANG TERPENGARUH

  fmt.Println("AFFECTED ROWS ", rows_affected);
}

func main() {
  var data_siswa []Siswa = getData(); // test ambil data dari database dan masukkan ke variabel
  var json, err = json.Marshal(data_siswa);
  if err != nil {
    fmt.Println(err);
    return
  }
  fmt.Println(string(json));
  
  //~ saveData(Siswa{Nama: "Nama Baru", Kelas: "Kelas Baru"});
  //~ updateData(Siswa{Id: 12, Nama: "Nama Baru", Kelas: "Kelas Baru"});
  //~ deleteData("10");
}

