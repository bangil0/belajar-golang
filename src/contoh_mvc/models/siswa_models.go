package models

import "fmt"

// struct pengganti class pada OOP
// PRIVATE : HURUF AWALNYA HARUS KECIL
// PUBLIC : HURUF AWALNYA HARUS BESAR
type Siswa struct {
  nama, kelas string
  umur int
}

// menambahkan method ke struct

func (s Siswa) GetNama() (nama string) { // (s Siswa) menandakan method ini milik Siswa
  nama = fmt.Sprintf("Nama Saya %v \n", s.nama); // dan variabel s menjadi cara mengakses properti struct
  return;
}

func (s Siswa) GetKelas() (kelas string) { 
  kelas = fmt.Sprintf("Saya kelas %v \n", s.kelas);
  return;
}

func (s Siswa) GetUmur() (umur string) { 
  umur = fmt.Sprintf("Umur Saya %v \n", s.umur);
  return;
}

func (s *Siswa) SetNama(nama string) {
  s.nama = nama;
}

func (s *Siswa) SetKelas(kelas string) {
  s.kelas = kelas;
}

func (s *Siswa) SetUmur(umur int) {
  s.umur = umur;
}

