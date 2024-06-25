package sake

import "net/http"

type RouteHandlers struct {
	route   string
	handler http.HandlerFunc
}

func RegisterRoutes(mux *http.ServeMux, routeHandlers []RouteHandlers) {
	for _, rh := range routeHandlers {
		mux.HandleFunc(rh.route, rh.handler)
	}
}
