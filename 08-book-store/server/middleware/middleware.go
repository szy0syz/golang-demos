package middleware

import (
	"log"
	"mime"
	"net/http"
)

// Logging 感觉还是和nodejs有相通的地方 😁
func Logging(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, req *http.Request) {
		log.Printf("revc a %s request from %s", req.Method, req.RemoteAddr)
		next.ServeHTTP(w, req)
	})
}

func Validating(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, req *http.Request) {
		contentType := req.Header.Get("Content-Type")
		mediaType, _, err := mime.ParseMediaType(contentType)
		if err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
		}

		if mediaType != "application/json" {
			http.Error(w, "invalid Content-Type", http.StatusUnsupportedMediaType)
		}

		next.ServeHTTP(w, req)
	})
}
