package main

import (
  "fmt"
)

func main() {
  var segitiga = func(alas, tinggi int) (luas float32) {
    luas = float32(alas*tinggi)/2;
    return;
  }
  var segitiga1 = segitiga(12, 10);
  fmt.Println(segitiga1);
  
  var kubus = func(sisi int) (volume int) { 
    volume = sisi*sisi*sisi;
    return;
  }(10); // membuat dan menjalankan fungsi anonimous secara langsung
  fmt.Printf("Volume kubus %v ", kubus);
}
