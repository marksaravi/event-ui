package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"os"
)

var rootDir string = "/"

func index(w http.ResponseWriter, req *http.Request) {
	bytesData, err := ioutil.ReadFile(fmt.Sprintf("%s/public/index.html", rootDir))
	if err != nil {
		log.Fatal(err)
	}

	html := string(bytesData)
	fmt.Fprintf(w, html)
}

func appjs(w http.ResponseWriter, req *http.Request) {
	bytesData, err := ioutil.ReadFile(fmt.Sprintf("%s/public/static/app.js", rootDir))
	if err != nil {
		log.Fatal(err)
	}

	js := string(bytesData)
	fmt.Fprintf(w, js)
}

func main() {
	rootDir, _ = os.Getwd()
	http.HandleFunc("/", index)
	http.HandleFunc("/static/app.js", appjs)

	http.ListenAndServe(":8090", nil)
}
