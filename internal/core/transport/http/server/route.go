package core_http_server

import (
	"net/http"

	core_http_middleware "github.com/Mixone1337/todo-app/internal/core/transport/http/middleware"
)

type Route struct {
	Method     string // get post и т.д
	Path       string // /tasks/{id} и т.д
	Handler    http.HandlerFunc
	Middleware []core_http_middleware.Middleware
}

func (r *Route) WithMiddleware() http.Handler {
	return core_http_middleware.ChainMiddleware(r.Handler, r.Middleware...)
}
