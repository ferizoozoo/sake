package main

import (
	"fmt"
	"io"
	"net/http"
)

func main() {
	fmt.Println("Hello to Sake!")
	mux := http.NewServeMux()

	handlerFuncs := []RouteHandlers{
		{
			"/",
			func(w http.ResponseWriter, r *http.Request) {
				fmt.Println("Request to / received")
				io.WriteString(w, "Hello to Sake!")
			},
		},
		{
			"/welcome",
			func(w http.ResponseWriter, r *http.Request) {
				fmt.Println("Request to /welcome received")
				io.WriteString(w, "Welcome to Sake!")
			},
		},
	}
	RegisterRoutes(mux, handlerFuncs)
	_ = http.ListenAndServe(":8000", mux)
}
