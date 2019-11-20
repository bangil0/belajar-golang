package main

import (
  "fmt"
  "strings"
)

func gabungkanKata(kata []string, separator string) {
  fmt.Println(strings.Join(kata, separator)); // menggabungkan isi array jadi string menggunakan separator
}

func main() {
  gabungkanKata([]string{"Ayam", "Kaki", "Balang"}, "---");
}
