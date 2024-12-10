package middleware

import (
	"backend/internal/helpers"
	"context"
	"net/http"
	"strings"
)

func writeUnauthorized(w http.ResponseWriter) {
	w.WriteHeader(http.StatusUnauthorized)
	w.Write([]byte(http.StatusText(http.StatusUnauthorized)))
}

func IsAuthenticated(next http.HandlerFunc) http.HandlerFunc {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		auth := r.Header.Get("Authorization")

		if !strings.HasPrefix(auth, "Bearer") {
			writeUnauthorized(w)
			return
		}

		encodedToken := strings.TrimPrefix(auth, "Bearer ")

		claims, err := helpers.DecodeJwt(encodedToken)
		if err != nil {
			writeUnauthorized(w)
			return
		}

		sub := claims["sub"].(float64)
		ctx := context.WithValue(r.Context(), "userId", int32(sub))
		r = r.WithContext(ctx)

		next.ServeHTTP(w, r)
	})
}
