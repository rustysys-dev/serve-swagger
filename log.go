package main

import (
	"log"
	"net/http"
)

// AccessLogger ...
func AccessLogger(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		log.Printf("URL called: %s", r.URL)
		next.ServeHTTP(w, r)
	})
}
