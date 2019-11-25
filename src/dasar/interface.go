package main

import "fmt"
import "encoding/json"

// interface adalah rancangan dari struct. interface hanya bisa menampung method kosong.
// dan method kosong tersebut nantinya akan diterapkan ke struct

// interface kosong akan menjadi tipe data dinamis

// fungsi dengan parameter interface kosong akan bisa menerima berbagai macam nilai alias dinamis


func main() {
  // contoh variabel map dengan tipe interface
  // menjadikan isi setiap map tersebut bisa dinamis
  var person = []map[string]interface{}{
    {"name": "Wick", "age": 23},
    {"name": "Ethan", "age": 23},
    {"name": "Bourne", "age": 22},
  }
  
  // single map menggunakan interface
  var madam = map[string]interface{}{
    "nama": "Madam",
    "umur": "25",
  }
    
  var json_person, _ = json.Marshal(person);
  var json_madam, _ = json.Marshal(madam);
  for _, p := range person {
    fmt.Printf("Name: %v, umur: %v \n ", p["name"], p["age"]);
  }
  
  fmt.Println(string(json_person));
  fmt.Println(string(json_madam));
}
