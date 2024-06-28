package sake

import "net/http"

type Server struct {
	mux *http.ServeMux
	mc  *MiddlewareContainer
}

func NewServer() *Server {
	return &Server{
		mux: http.NewServeMux(),
		mc:  NewMiddlewareContainer(),
	}
}
