package main

import (
	"fmt"
	"net/http"
	"os"
)

func main() {
	url := os.Args[1]
	tre := url[:3]

	if tre != "htt" {
		url = "http://" + url
	}

	fmt.Println(url)
	resp, err := http.Get(url)
	if (err == nil) && (resp.StatusCode == 200) {
		fmt.Println("UP")
	} else {
		fmt.Println("DOWN")
	}
}
