package main

import (
  "fmt"
  "net/http"
  "os"
  "io/ioutil"
)

func main(){
  var url string
  fmt.Println(">> url: (Make sure to use http://)")
  fmt.Scanf("%s", &url)
  data, err := http.Get(url)
  if err != nil{
    os.Exit(1)
  }else{
    defer data.Body.Close()
    contents, err := ioutil.ReadAll(data.Body)
    if err != nil{
      fmt.Println("%s\n", err)
    }
    fmt.Println("%s\n", string(contents))
  }
}
