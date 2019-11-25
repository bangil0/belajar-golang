package main

import "fmt"
import "net/http" // package untuk web server
import "html/template" // package untuk template engine

// fungsi yang akan dijalankan pada saat route tertentu diakses
// mirip controller
// parameternya wajib seperti ini
func contohHalaman(res http.ResponseWriter, req *http.Request) {
    fmt.Fprintln(res, "apa kabar!")
}

func contohHalamanPakaiTemplate(res http.ResponseWriter, req *http.Request) {
  // data yang akan dimasukkan kedalam html
  var data = map[string]interface{}{
    "Hobi": []string{"Satu", "Dua","tiga"},
    "Nama": "Madam",
    "Umur": 24,
  }
  
  // proses parse data ke template
  var t, err = template.ParseFiles("html/madam.html")
  
  // proses pengecekan error
  if err != nil {
      fmt.Println(err.Error())
      return
  }

  t.Execute(res, data)
}

func main() {
  
    // contoh route
    http.HandleFunc("/", func(res http.ResponseWriter, req *http.Request) {
        fmt.Fprintln(res, "halo!")
    })
    
    
    // contoh route
    http.HandleFunc("/index", contohHalaman) 
    
    // contoh route
    http.HandleFunc("/madam", contohHalamanPakaiTemplate) 

    fmt.Println("starting web server at http://localhost:8080/")
    http.ListenAndServe(":8080", nil)
}
