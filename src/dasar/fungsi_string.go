package main

import (
  "strings"
  "fmt"
)

func main() {
  var hasil = strings.Contains("anu madam", "madam"); // mengecek apakah tulisan madam ada didalam tulisan anu madam. hasil return bool
  var isPrefix1 = strings.HasPrefix("john wick", "jo"); // mengecek apakah tulisan jo ada diawal tulisan john wich atau fungsi untuk mengecek prefix. hasil return bool.
  var isSuffix1 = strings.HasSuffix("john wick", "ic"); // kebalikan dari prefix, yaitu cek tulisan dari belakang
  var howMany = strings.Count("ethan hunt", "t"); // menghitung banyak huruf t di tulisan ethan hunt. hasil return int
  var index1 = strings.Index("ethan hunt", "ha"); // mencari posisi tulisan ha di ethan hunt. yaitu diposisi atau index ke 2 dari sebelah kiri. hasil return int
  var newText1 = strings.Replace("anu madam", "m", "b", 1); // mengganti huruf M di tulisan "anu madam" jadi huruf B sebanyak 1 kali. ubah 1 jadi -1 agar semua huruf diganti
  var str = strings.Repeat("na", 4); // menampilkan tulisan na sebanyak 4 kali
  var string1 = strings.Split("the dark knight", " "); // mirip explode di PHP. ganti spasi jadi kutip kosong agar semua huruf jadi array
  var data = []string{"banana", "papaya", "tomato"};
  var str = strings.Join(data, "-"); // implode pada PHP
  var str = strings.ToLower("aLAy"); // mengubah semua huruf jadi huruf kecil
  var str = strings.ToUpper("eat!"); // mengubah semua huruf jadi huruf besar
   
}
