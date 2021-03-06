package main

import "fmt"
import "net/http" // package untuk web server
import "html/template" // package untuk template engine

// contoh penggunaan sistem template/layout 

func LoadView(data map[string]interface{}, content string, templateName string, res http.ResponseWriter) {
  // tentukan terlebih dahulu html mana saja yang akan dipakai
  var layout = template.Must(template.ParseFiles(
       "layout_html/bagian_head.html",
       "layout_html/bagian_header.html",
       content,
       "layout_html/bagian_footer.html",
   ))
   
  // cek file layout_html/halaman_contoh.html untuk mengecek nama templatenya dibagian DEFINE
  var err = layout.ExecuteTemplate(res, templateName, data)
  
  // cek error
  if err != nil {
      http.Error(res, err.Error(), http.StatusInternalServerError)
  }
}

func contoh(res http.ResponseWriter, req *http.Request) {
  // data yang akan dimasukkan kedalam html
  var data = map[string]interface{}{
    "Hobi": []string{"Satu", "Dua","tiga"},
    "Nama": "Madam",
    "Umur": 24,
  }
  
  LoadView(data, "layout_html/halaman_contoh.html", "halaman_contoh", res);
}

func main() {

  // contoh route
  http.HandleFunc("/", contoh) 

  fmt.Println("starting web server at http://localhost:8080/")
  http.ListenAndServe(":8080", nil)
}
