package core_http_server

import "net/http"

type Route struct {
	Method  string // get post и т.д
	Path    string // /tasks/{id} и т.д
	Handler http.HandlerFunc
}

func NewRoute(
	method string,
	path string,
	handler http.HandlerFunc,
) Route {
	return Route{
		Method:  method,
		Path:    path,
		Handler: handler,
	}
}
