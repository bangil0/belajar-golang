package controllers

import (
  "contoh_mvc/models"
  "contoh_mvc/helper"
  "net/http"
  "strconv"
)

type Message struct {
  Code int `json:code`
  Message string `json:message`
  Data string `json:data`
}

var siswa models.Siswa;
var message Message;

func ListSiswa(res http.ResponseWriter, req *http.Request){
  // untuk menentukan header response
  var data = map[string]interface{}{
    "DataSiswa": siswa.SelectAll(),
  }
  
  helper.LoadView(data, "views/page/data_siswa.html", "data_siswa", res);
}

func EditSiswa(res http.ResponseWriter, req *http.Request){
  id, err := strconv.Atoi(req.FormValue("id_siswa")); // konversi string ke integer
  if err == nil {
    if req.Method == "POST" {
      siswa.SetId(id);
      siswa.SetNama(req.FormValue("nama"));
      siswa.SetKelas(req.FormValue("kelas"));
      siswa.Update();
      http.Redirect(res, req, "/siswa", http.StatusSeeOther)
    } else {
      var id_siswa string = req.FormValue("id_siswa"); 
      // untuk menentukan header response
      var data = map[string]interface{}{
        "DataSiswa": siswa.Select(id_siswa),
      }
      
      helper.LoadView(data, "views/page/edit_siswa.html", "edit_siswa", res);
    }
  }  
}

func AddSiswa(res http.ResponseWriter, req *http.Request){
  if req.Method == "GET" {
    helper.LoadView(nil, "views/page/add_siswa.html", "add_siswa", res);  
  } else {
    siswa.SetNama(req.FormValue("nama"));
    siswa.SetKelas(req.FormValue("kelas"));
    siswa.Insert();
    http.Redirect(res, req, "/siswa", http.StatusSeeOther)
  }
  
}
func DeleteSiswa(res http.ResponseWriter, req *http.Request){
  id, err := strconv.Atoi(req.FormValue("id_siswa")); // konversi string ke integer
  if err == nil {
    siswa.SetId(id);
    siswa.Delete();
  }
  http.Redirect(res, req, "/siswa", http.StatusSeeOther)
}




