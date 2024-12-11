package middleware

import (
	"fmt"
	"net/http"
	"time"
)

func Logger(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		fmt.Printf("%s %s %s %v \n", r.RemoteAddr, r.Method, r.URL, time.Now().UTC().Format(time.RFC3339))
		next.ServeHTTP(w, r)
	})
}
