package main

import (
	"fmt"
	"net/http"
)

func main() {
	port := 3000
	fmt.Println(fmt.Sprintf("Starting Server on Port: %d", port))

	http.HandleFunc("/", routes)
	http.ListenAndServe(fmt.Sprintf(":%d", port), nil)
}

func routes(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Hello, World"))
}
