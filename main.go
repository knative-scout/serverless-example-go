package main

import (
	"fmt"
	"html"
	"log"
	"net/http"
)

func main() {
	http.HandleFunc("/", request_handler)
	fmt.Println("Server started ..")
	log.Fatal(http.ListenAndServe(":8080", nil))
}

func request_handler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Hello, %q", html.EscapeString(r.URL.Path)[1:])
}
