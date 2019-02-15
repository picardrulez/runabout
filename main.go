package main

import (
	"flag"
	"fmt"
	"net/http"
	"os"
)

var VERSION = "v1.1"

func main() {
	versionFlag := flag.Bool("v", false, "display version")
	flag.Parse()
	if *versionFlag {
		fmt.Println("runabout " + VERSION)
		return
	}
	url := os.Args[1]
	tre := url[:3]

	if tre != "htt" {
		url = "http://" + url
	}

	resp, err := http.Get(url)
	if (err == nil) && ((resp.StatusCode == 200) || (resp.StatusCode == 302)) {
		fmt.Println("UP")
	} else {
		fmt.Println("DOWN")
	}
}
