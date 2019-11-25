package main

import "fmt"
import "net/url" // package untuk handle parsing url

func main() {
  var urlString = "http://kalipare.com:80/hello?name=john wick&age=27"
  var u, e = url.Parse(urlString) // method untuk parse url
  
  // cek proses parsing apakah ada error atau tidak
  if e != nil {
      fmt.Println(e.Error())
      return
  }
  
  // menampilkan full alamat mentah
  fmt.Printf("url: %s\n", urlString)
  
  fmt.Printf("protocol: %s\n", u.Scheme) // http
  fmt.Printf("host: %s\n", u.Host)       // kalipare.com:80
  fmt.Printf("path: %s\n", u.Path)       // /hello
  
  // bagian query string
  var name = u.Query()["name"][0] // john wick
  var age = u.Query()["age"][0]   // 27
  fmt.Printf("name: %s, age: %s\n", name, age)
}
