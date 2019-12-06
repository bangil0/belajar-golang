package controllers

import (
  "contoh_mvc/models"
  "contoh_mvc/helper"
  "net/http"
  "strconv"
)
type SiswaController struct {
  SiswaModel models.Siswa;
}

// method untuk inisialisasi struct SiswaController, design pattern factory
func NewSiswaController() SiswaController {
  return SiswaController{}
}

func (this SiswaController) ListSiswa(res http.ResponseWriter, req *http.Request){
  // untuk menentukan header response
  var data = map[string]interface{}{
    "DataSiswa": this.SiswaModel.SelectAll(),
  }
  
  helper.LoadView(data, "views/page/data_siswa.html", "data_siswa", res);
}

func (this SiswaController) EditSiswa(res http.ResponseWriter, req *http.Request){
  id, err := strconv.Atoi(req.FormValue("id_siswa")); // konversi string ke integer
  if err == nil {
    if req.Method == "POST" {
      this.SiswaModel.SetId(id);
      this.SiswaModel.SetNama(req.FormValue("nama"));
      this.SiswaModel.SetKelas(req.FormValue("kelas"));
      this.SiswaModel.Update();
      http.Redirect(res, req, "/siswa", http.StatusSeeOther)
    } else {
      var id_siswa string = req.FormValue("id_siswa"); 
      // untuk menentukan header response
      var data = map[string]interface{}{
        "DataSiswa": this.SiswaModel.Select(id_siswa),
      }
      
      helper.LoadView(data, "views/page/edit_siswa.html", "edit_siswa", res);
    }
  }  
}

func (this SiswaController) AddSiswa(res http.ResponseWriter, req *http.Request){
  if req.Method == "GET" {
    helper.LoadView(nil, "views/page/add_siswa.html", "add_siswa", res);  
  } else {
    this.SiswaModel.SetNama(req.FormValue("nama"));
    this.SiswaModel.SetKelas(req.FormValue("kelas"));
    this.SiswaModel.Insert();
    http.Redirect(res, req, "/siswa", http.StatusSeeOther)
  }
}
func (this SiswaController) DeleteSiswa(res http.ResponseWriter, req *http.Request){
  id, err := strconv.Atoi(req.FormValue("id_siswa")); // konversi string ke integer
  if err == nil {
    this.SiswaModel.SetId(id);
    this.SiswaModel.Delete();
  }
  http.Redirect(res, req, "/siswa", http.StatusSeeOther)
}




