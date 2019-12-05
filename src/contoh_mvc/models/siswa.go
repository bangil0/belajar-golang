package models

import (
  "fmt"
  "database/sql" // datase
  _ "github.com/go-sql-driver/mysql" // import package db
)

type Siswa struct {
  Id int `json:"id"`
  Nama string `json:"nama"`
  Kelas string `json:"kelas"`
}

func db() (*sql.DB, error) {
  db, err := sql.Open("mysql", "root:mysql@tcp(localhost:3306)/golang") // terhubung ke database
  if err != nil { // cek error
      return nil, err
  }
  return db, nil // kembalikan instansi database
}

// menambahkan method ke struct
func (s Siswa) GetNama() (nama string) { // (s Siswa) menandakan method ini milik Siswa
  nama = fmt.Sprintf("Nama Saya %v \n", s.Nama); // dan variabel s menjadi cara mengakses properti struct
  return;
}

func (s Siswa) GetKelas() (kelas string) { 
  kelas = fmt.Sprintf("Saya kelas %v \n", s.Kelas);
  return;
}

func (s *Siswa) SetId(id int) {
  s.Id = id;
}

func (s *Siswa) SetNama(nama string) {
  s.Nama = nama;
}

func (s *Siswa) SetKelas(kelas string) {
  s.Kelas = kelas;
}

func (s *Siswa) SelectAll() (data_siswa []Siswa) {
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

func (s *Siswa) Select(id string) (data_siswa Siswa) {
  db, err := db() // ambil koneksi database ke variabel
  if err != nil { // cek error
    fmt.Println(err.Error())
    return
  }
  defer db.Close() // tutup koneksi nanti 
  
  err = db.QueryRow("select * from siswa where id = ?", id).Scan(&data_siswa.Id, &data_siswa.Nama, &data_siswa.Kelas) // jalankan query sql
  if err != nil { // cek eksekusi error
      fmt.Println(err.Error())
      return
  }
  
  return
}

func (data_siswa *Siswa) Insert() {
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

func (data_siswa *Siswa) Update() {
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
func (data_siswa *Siswa) Delete() {
  db, err := db() // ambil koneksi database ke variabel
  if err != nil { // cek error
    fmt.Println(err.Error())
    return
  }
  defer db.Close() // tutup koneksi nanti 
  
  hasil, err := db.Exec("DELETE FROM siswa WHERE id = ?", data_siswa.Id);
  if err != nil {
    fmt.Println(err);
    return
  }
  rows_affected, _ := hasil.RowsAffected(); // AMBIL BARIS YANG TERPENGARUH

  fmt.Println("AFFECTED ROWS ", rows_affected);
}

