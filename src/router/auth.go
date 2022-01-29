package router

import (
	"net/http"
	"strings"
)

func auth(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		authHeader := strings.Split(r.Header.Get("Authorization"), "Bearer ")
		if len(authHeader) >= 2 {
			for _, key := range config.Keys {
				if key == authHeader[1] {
					next.ServeHTTP(w, r)
					return
				}
			}
		}
		w.Write([]byte("Missing authorization"))
		return
	})
}
