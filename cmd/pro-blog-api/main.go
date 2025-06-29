package main

import (
	"fmt"
	"net/http"
)

func main() {
	http.HandleFunc("/home", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintln(w, "Home Page")
	})

	fmt.Println("Server started on :8080")
	http.ListenAndServe(":8080", nil)
}
