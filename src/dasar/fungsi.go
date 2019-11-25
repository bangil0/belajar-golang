package main

import (
  "fmt"
  "math"
)

// contoh fungsi dengan parameter variadic atau satu parameter banyak isi
func rataRata(daftar_nilai ...float32) (float32) {
  var total float32; // buat penampung total
  var banyak_nilai float32 = float32(len(daftar_nilai)); // hitung banyak nilai
  for _, nilai := range daftar_nilai {
    total += nilai;
  }
  return total/banyak_nilai;
}

// contoh fungsi biasa
func hitungLuasSegitiga(alas, tinggi float32) float32 {
  var luas float32 = ( alas * tinggi ) / 2;
  return luas;
}

// contoh fungsi yang mengembalikan banyak nilai
func balikanParameterFungsi(huruf1, huruf2 string) (string, string) {
  return huruf2, huruf1;
}

// menghitung volume kubus
func volumeKubus(sisi float64) float64 {
  var volume float64 = math.Pow(sisi, 3);
  return volume;
}

func luasPersegi(sisi  float64) float64 {
  return math.Pow(sisi, 2);
}

func main() {
  // memasukkan fungsi banyak nilai ke variabel
  var hasil1, hasil2 string = balikanParameterFungsi("Huruf Awal", "Huruf AKhir");
  var luas_segitiga float32 = hitungLuasSegitiga(10.19, 20.21);
  var volume_kubus float32 = float32(volumeKubus(12));
  var luas_persegi float32 = float32(luasPersegi(12));
  var nilai_siswa float32 = rataRata(1, 2, 3, 4, 5, 6, 7);
  
  /*
   * format string pada printf :
   * %v = format bawaan. cocok untuk kasus umum
   *  
   */
	fmt.Printf("Luas segitiga adalah %v cm \n", luas_segitiga);
	fmt.Printf("Volume kubus adalah %v cm \n", volume_kubus);
	fmt.Printf("Luas persegi adalah %v cm \n", luas_persegi);
  fmt.Printf("Contoh fungsi yang mengembalikan 2 nilai : nilai pertama %v dan nilai kedua %v \n", hasil1, hasil2);
  fmt.Printf("Contoh fungsi variadic atau satu parameter banyak isi seperti array. nilai rata-ratanya adalah %v \n", nilai_siswa);
}
