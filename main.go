package main

import (
	"fmt"
	"io"
	"net/http"
)

func main() {
	fmt.Println("Hello to Sake!")
	mux := http.NewServeMux()
	mux.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Println("Request to / received")
		io.WriteString(w, "Hello to Sake!")
	})
	_ = http.ListenAndServe(":8000", mux)
}
