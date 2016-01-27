package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"strings"
)

// askURL provide an i/o process asking for a valid URL
func askURL() string {
	var url string
	fmt.Print("URL: ")
	fmt.Scanf("%s", &url)

	if (url == "") ||
		((!strings.HasPrefix(url, "http:")) ||
			(!strings.HasPrefix(url, "http:"))) {
		fmt.Println("Please type in a valid url")
		url = askURL()
	}

	return url
}

func main() {
	fmt.Println(">> Insert url: (Make sure to use http://)")
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
