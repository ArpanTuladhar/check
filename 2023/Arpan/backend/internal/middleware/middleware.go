package middleware

import (
	"net/http"
)

type Middleware interface {
	WithAuth() func(http.Handler) http.Handler
}

type middleware struct {
}

func NewMiddleware() Middleware {
	return middleware{}
}
