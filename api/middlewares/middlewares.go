package middlewares

import (
	"errors"
	"net/http"

	"github.com/nstoker/apiance1/api/auth"
	"github.com/nstoker/apiance1/api/responses"
)

// SetMiddlewareJSON sets the middleware json
func SetMiddlewareJSON(next http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		next(w, r)
	}
}

// SetMiddlewareAuthentication sets the middleware authentication
func SetMiddlewareAuthentication(next http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		err := auth.TokenValid(r)
		if err != nil {
			responses.ERROR(w, http.StatusUnauthorized, errors.New("Not Authorized (SMA)"))
			return
		}
		next(w, r)
	}
}
