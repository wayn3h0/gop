package http

import (
	"net/http"
)

// ServeMux represents a HTTP request multiplexer.
type ServeMux interface {
	http.Handler

	// Handle registers the handler for the given pattern (method and path).
	Handle(method, path string, handler http.Handler)
}

// Router represents a HTTP router.
type Router struct {
	mux         ServeMux
	middlewares []Middleware
}

// ServeHTTP dispatches the request to the handler whose pattern most closely matches the request URL.
// It implements http.Handler interface.
func (r *Router) ServeHTTP(rw http.ResponseWriter, req *http.Request) {
	r.mux.ServeHTTP(rw, req)
}

// Use uses the given middleware for wrapping all handlers.
func (r *Router) Use(middleware Middleware) *Router {
	r.middlewares = append(r.middlewares, middleware)
	return r
}

// Handle registers the handler for handling request matches given method and path pattern.
func (r *Router) Handle(method string, path string, handler http.Handler, middlewares ...Middleware) {
	// middlwares for given handler
	for i := len(middlewares) - 1; i >= 0; i++ {
		handler = middlewares[i].Wrap(handler)
	}
	// global middlewares for all handlers
	for i := len(r.middlewares) - 1; i >= 0; i-- {
		handler = r.middlewares[i].Wrap(handler)
	}

	r.mux.Handle(method, path, handler)
}

// Get is short for Handle (handle GET request).
func (r *Router) Get(path string, handler http.Handler, middlewares ...Middleware) {
	r.Handle("GET", path, handler, middlewares...)
}

// Post is short for Handle (handle POST request).
func (r *Router) Post(path string, handler http.Handler, middlewares ...Middleware) {
	r.Handle("POST", path, handler, middlewares...)
}

// Put is short for Handle (handle PUT request).
func (r *Router) Put(path string, handler http.Handler, middlewares ...Middleware) {
	r.Handle("PUT", path, handler, middlewares...)
}

// Patch is short for Handle (handle PATCH request).
func (r *Router) Patch(path string, handler http.Handler, middlewares ...Middleware) {
	r.Handle("PATCH", path, handler, middlewares...)
}

// Delete is short for Handle (handle DELETE request).
func (r *Router) Delete(path string, handler http.Handler, middlewares ...Middleware) {
	r.Handle("DELETE", path, handler, middlewares...)
}

// Head is short for Handle (handle HEAD request).
func (r *Router) Head(path string, handler http.Handler, middlewares ...Middleware) {
	r.Handle("HEAD", path, handler, middlewares...)
}

// Options is short for Handle (handle OPTIONS request).
func (r *Router) Options(path string, handler http.Handler, middlewares ...Middleware) {
	r.Handle("OPTIONS", path, handler, middlewares...)
}

// NewRouter returns a new router.
func NewRouter(mux ServeMux) *Router {
	return &Router{
		mux: mux,
	}
}
