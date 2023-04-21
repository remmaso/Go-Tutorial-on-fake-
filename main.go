package main

import (
	"fmt"
	"net/http"
)

func main() {
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "Hi Ralph, You are a good developer!")
	})

	http.ListenAndServe(":8080", nil)
}
