package main

import (
	"fmt"
	"net/http"
)

func main() {
	// Use http.FileServer to serve static files from the "public" directory
	fs := http.FileServer(http.Dir("public"))
	http.Handle("/", fs)

	fmt.Println("Server listening on port 3000")
	http.ListenAndServe(":3000", nil)
}