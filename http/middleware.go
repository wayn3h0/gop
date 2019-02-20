package http

import (
	"net/http"
)

// Middleware represents a wrapper for http handler.
type Middleware interface {
	// Wrap waprs the handler.
	Wrap(http.Handler) http.Handler
}

// MiddlewareFunc an adapter to allow the use of ordinary functions as middlewares.
type MiddlewareFunc func(http.Handler) http.Handler

// Wrap warps the m to http.Handler.
// Wrap implements Middleware interface.
func (m MiddlewareFunc) Wrap(h http.Handler) http.Handler {
	return m(h)
}
