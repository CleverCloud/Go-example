package main

import (
	"fmt"
	"net/http"
)

func main() {
	http.HandleFunc("/", indexPage)
	http.ListenAndServe("0.0.0.0:8080", nil)
}

func indexPage(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(
		w, `<h1>Hello, World</h1>
    <p>Go is running on Clever Cloud ๐กโ๏ธ,</p>
    <p>you should give it a try too!</p>`)

}

//do anything
