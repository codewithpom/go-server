package main

import (
	"fmt"
	"net/http"
)

func greet(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Hello, World!")

}

func main() {
	http.HandleFunc("/", greet)
	http.ListenAndServe(":8080", nil)
	fmt.Println("Server started")
}
