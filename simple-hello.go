package main

import (
	"fmt"
	"net/http"
)

func handler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Hello world from my Go program!")
}

func main() {
// use envconfig to put environment variables into the HTTP output to show it works
	http.HandleFunc("/", handler)              // redirect all urls to the handler function
	http.ListenAndServe("0.0.0.0:9999", nil) // listen for connections at port 9999 on the local machine
}
