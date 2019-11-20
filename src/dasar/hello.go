package main

import "fmt"

func main() {
  var tulisan string = fmt.Sprintf("Sprintf sama seperti printf, \n tapi hasilnya bisa disimpan ke variabel %v \n", 123);
	fmt.Printf("Luas segitiga adalah %v cm \n", 123);
  fmt.Print(tulisan);
}
