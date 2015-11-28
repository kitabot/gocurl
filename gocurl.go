// gocurl is a simple curl like program that able to fetch site's index.html page
// under some massive development
// MIT - DWTFPL

package main

import (
  "fmt"
  "net/http"
  "os"
  "io/ioutil"
)

  func main(){
  var url string
  fmt.Println(">> url: ")
  fmt.Scanf("%s", &url)
  if url == ""{
    fmt.Println("unable to fetch any data")
    os.Exit(1)
  }
  fmt.Println("fetching ",url)
  data, err := http.Get(url)

  if err != nil{
    fmt.Println("something went wrong...")
    fmt.Println(err)
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
