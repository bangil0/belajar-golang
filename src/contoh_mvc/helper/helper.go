package helper

import (
  "html/template"
  "net/http"
)

func LoadView(data map[string]interface{}, content string, templateName string, res http.ResponseWriter) {
  // tentukan terlebih dahulu html mana saja yang akan dipakai
  var layout = template.Must(template.ParseFiles(
       "views/layout/bagian_head.html",
       "views/layout/bagian_header.html",
       content,
       "views/layout/bagian_footer.html",
   ))
   
  // cek file layout_html/halaman_contoh.html untuk mengecek nama templatenya dibagian DEFINE
  var err = layout.ExecuteTemplate(res, templateName, data)
  
  // cek error
  if err != nil {
      http.Error(res, err.Error(), http.StatusInternalServerError)
  }
}
