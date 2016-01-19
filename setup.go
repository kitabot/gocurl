package main

import (
  "fmt"
  "net/http"
  "log"
  "os"
  "io/ioutil"
)

func main(){
  var url string
  fmt.Println(">> Insert url: (Make sure to use http://)")
  fmt.Scanf("%s", &url)
  if url == "" {
    fmt.Println("Please type in an url")
    os.Exit(1)
  }
  data, err := http.Get(url)
  if err != nil{
    log.Error(err)
    os.Exit(1)
  }else{
    fmt.Println("Fetching url...")
    defer data.Body.Close()
    contents, err := ioutil.ReadAll(data.Body)
    if err != nil{
      log.Error(err)
    }
    fmt.Println("%s\n", string(contents))
  }
}
