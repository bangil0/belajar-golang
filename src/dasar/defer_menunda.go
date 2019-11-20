package main

import "fmt"

func main() {
  defer fmt.Println("Ini akan keluar dibagian terakhir. tulisan ini akan muncul saat sebuah fungsi telah sampai di return atau diakhir baris fungsi");
  defer fmt.Println("terakhir 2");
  defer fmt.Println("terakhir 3");
  fmt.Println("1");
  fmt.Println("2");
  fmt.Println("3");
  fmt.Println("4");
}
