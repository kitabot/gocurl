package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"strings"
	"os"
)

// askURL provide an i/o process asking for a valid URL
func askURL() string {
	var url string
	fmt.Print("URL: ")
	fmt.Scanf("%s", &url)

	if (url == "exit") {
		os.Exit(1)
	}

	if (url == "") ||
		((!strings.HasPrefix(url, "http:")) ||
			(!strings.HasPrefix(url, "https:"))) {
		fmt.Println("Please type in a valid url or use http://")
		url = askURL()
	}	

	return url
}

func main() {	
	url := askURL()

	data, err := http.Get(url)
	if err != nil {
		log.Fatal(err)
	} else {
		fmt.Println("Fetching url...")
		defer data.Body.Close()
		contents, err := ioutil.ReadAll(data.Body)
		if err != nil {
			log.Fatal(err)
		}
		fmt.Printf("%s\n", string(contents))
	}
}