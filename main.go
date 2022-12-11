package main

import (
	"fmt"
	"io"
	"net/http"
)

func main() {
	http.HandleFunc("/", index)

	fmt.Printf("initiating server on port 8080...")
	http.ListenAndServe(":8080", nil)
}

func index(w http.ResponseWriter, r *http.Request) {

	s := `<h1>Welcome to a server. In here you can serve files, pages, and more.</h1>
    <h2>I am Melvin Raj, writer of <a href = "https://getsetgo.substack.com" target="_blank" >Get Set Go</a></h2>
`
	io.WriteString(w, s)
}
