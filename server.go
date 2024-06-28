package sake

import "net/http"

type Server struct {
	mux *http.ServeMux
	//TODO: add a middleware list for using before any request
	// middlewares []Middleware
}

func NewServer() *Server {
	mux := http.NewServeMux()

	return &Server{
		mux,
	}
}
