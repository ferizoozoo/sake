package sake

import (
	"fmt"
	"io"
	"net/http"
	"testing"
)

const (
	host = "localhost"
	port = "8000"
)

var url string = fmt.Sprintf("http://%s:%s", host, port)

// TestHelloName calls greetings.Hello with a name, checking
// for a valid return value.
func TestRegisteredRouters(t *testing.T) {
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

	go getHttpMuxAndRegisterWithHandlers(handlerFuncs)

	for _, handlerFunc := range handlerFuncs {
		route := url + handlerFunc.route
		res, err := http.Get(route)
		if err != nil {
			t.Fatal(err)
		}

		if res.StatusCode != 200 {
			t.Fatalf("Error happened")
		}
	}
}

func getHttpMuxAndRegisterWithHandlers(routeHandlers []RouteHandlers) {
	mux := http.NewServeMux()
	RegisterRoutes(mux, routeHandlers)
	u := fmt.Sprintf("%s:%s", host, port)
	_ = http.ListenAndServe(u, mux)
}
