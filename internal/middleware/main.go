package middleware

import "net/http"

type MiddleWare func(http.Handler) http.Handler

func LoadMiddlewares(middlewares ...MiddleWare) MiddleWare {
	return func(next http.Handler) http.Handler {
		for i := len(middlewares) - 1; i >= 0; i-- {
			middleware := middlewares[i]
			next = middleware(next)
		}

		return next
	}
}
