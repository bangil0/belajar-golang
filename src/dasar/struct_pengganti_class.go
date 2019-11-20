package main

import "fmt"

// struct pengganti class pada OOP
type Siswa struct {
  nama, kelas string
  umur int
}

// menambahkan method ke struct

func (s Siswa) getNama() (nama string) { // (s Siswa) menandakan method ini milik Siswa
  nama = fmt.Sprintf("Nama Saya %v \n", s.nama); // dan variabel s menjadi cara mengakses properti struct
  return;
}

func (s Siswa) getKelas() (kelas string) { 
  kelas = fmt.Sprintf("Saya kelas %v \n", s.kelas);
  return;
}

func (s Siswa) setNama(nama string) {  // properti struct tidak bisa diubah lewat method dan hanya berpengaruh pada struct yang dimethod tersebut
  s.nama = nama;
}

func (s *Siswa) setKelas(kelas string) { // jadi kalau ingin merubah nilai propertinya, wajib pakai pointer atau bintang dinama structnya seperti yg disamping
  s.kelas = kelas;
}

func main() {
  var madam = Siswa{"Madam Retno", "XII-1", 12};
  var madam2 = struct { // penulisan struct secara langsung di variabel
      nama string
      kelas string
  }{
      nama: "Madam",
      kelas: "satu",
  };
  
  fmt.Println(madam2.nama);
  fmt.Println(madam2.kelas);
  
  fmt.Println(madam.getNama());
  fmt.Println(madam.getKelas());
  
  madam.setNama("Madam Baru"); // yang ini tidak berpengaruh karena tidak pakai pointer
  fmt.Println(madam.getNama());
  
  madam.setKelas("XII Baru"); // yang berpengaruh karena tidak pakai pointer
  fmt.Println(madam.getKelas());
}
