package controllers

import (
  "fmt"
  "contoh_mvc/models"
  "contoh_mvc/helper"
  "net/http"
  "encoding/json"
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
    "DataSiswa": siswa.GetData(),
  }
  
  helper.LoadView(data, "views/page/data_siswa.html", "data_siswa", res);
}
func GetSiswa(res http.ResponseWriter, req *http.Request){
  // untuk menentukan header response
  res.Header().Set("Content-Type", "application/json");
  var response, _ = json.Marshal(struct {
    Code int
    Message string
    Data []models.Siswa
  }{
    Code: 200,
    Message: "Ok",
    Data: siswa.GetData(),
  });
  fmt.Fprintln(res, string(response));
}

func SaveSiswa(res http.ResponseWriter, req *http.Request){
  // untuk menentukan header response
  res.Header().Set("Content-Type", "application/json");
  
  if req.Method == "POST" {
    siswa.SetNama(req.FormValue("nama"));
    siswa.SetKelas(req.FormValue("kelas"));
    siswa.SaveData();
    var response, _ = json.Marshal(struct {
      Code int
      Message string
      Data []models.Siswa
    }{
      Code: 200,
      Message: "Ok",
    });
    fmt.Fprintln(res, string(response));
  } else {
    var response, _ = json.Marshal(struct {
      Code int
      Message string
    }{
      Code: 404,
      Message: "Not Found",
    });
    fmt.Fprintln(res, string(response));
  }
}


