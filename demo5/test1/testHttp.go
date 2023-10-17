package main

import (
	"fmt"
	"net/http"
)

func welcome(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "text/html;charset=utf-8")
	fmt.Fprintln(w, "<b>欢迎光临我的网站</b>")

}
func main() {

	http.HandleFunc("/", welcome)
	http.ListenAndServe("localhost:8081", nil)
	fmt.Print("print success!")

}
