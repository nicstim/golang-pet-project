package main

import (
	"fmt"
	"net/http"
)

func index(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Method: "+r.Method)
}

func main() {
	http.HandleFunc("/", index)
	http.ListenAndServe(":4000", nil)
}
