package middleware

import (
	"backend/internal/helpers"
	"context"
	"fmt"
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
		fmt.Println(auth)

		if !strings.HasPrefix(auth, "Bearer") {
			writeUnauthorized(w)
			return
		}

		encodedToken := strings.TrimPrefix(auth, "Bearer ")

		claims, err := helpers.DecodeJwt(encodedToken)
		if err != nil {
			fmt.Println(err)
			writeUnauthorized(w)
			return
		}

		sub := claims["sub"]
		ctx := context.WithValue(r.Context(), "userId", sub)
		r = r.WithContext(ctx)

		next.ServeHTTP(w, r)
	})
}
