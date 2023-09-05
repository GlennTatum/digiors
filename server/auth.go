package server

import "net/http"

func authenticationHandler(w http.ResponseWriter, r *http.Request) bool {
	return true
}

func authentication(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		ok := authenticationHandler(w, r)
		if ok {
			next.ServeHTTP(w, r)
		}
	})
}
