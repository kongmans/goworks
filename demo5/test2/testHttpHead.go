package main

import (
	"fmt"
	"net/http"
)

func param(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "text/html;charset=utf-8")
	fmt.Fprintln(w, "<b>test2</b>")

}

func main() {

	http.HandleFunc("/", param)
	http.ListenAndServe("localhost:8082", nil)
	fmt.Print("print success!")

}
