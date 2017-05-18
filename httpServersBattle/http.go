package main

import (
	"net/http"
	"fmt"
)

func handleRootRequest(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "Hello, World!")
}

func main() {
	http.HandleFunc("/", handleRootRequest)
	fmt.Println(http.ListenAndServe(":7000", nil))
}