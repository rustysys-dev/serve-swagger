package main

import (
	"fmt"
	"net/http"

	"github.com/go-chi/chi/v5"
)

func main() {
	fmt.Println("hello world")
	r := chi.NewMux()
	r.Use(AccessLogger)
	r.Get("/swagger/*", SwaggerAPI().ServeHTTP)

	http.ListenAndServe(":3030", r)
}
