package main

import (
	"fmt"
	"net/http"
)

func handler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Wise Word")
}

func main() {
	http.HandleFunc("/", handler)
	fmt.Println("Someone Hit me!")
	http.ListenAndServe(":8080", nil)
}
