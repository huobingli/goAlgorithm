package main

import (
	"fmt"
	"net/http"

	fileoper "../goserver/fileoper"
)

func IndexHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "world")
}

func HelloHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "hello")
}

func main() {
	// http.HandleFunc("/index", IndexHandler)

	// http.HandleFunc("/hello", HelloHandler)
	// fmt.Println(http.ListenAndServe("0.0.0.0:8090", nil))

	fileoper.Openfile()
}
