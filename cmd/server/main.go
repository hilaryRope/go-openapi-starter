package main

import (
	"github.com/go-chi/chi/v5"
	"go-openapi-starter/internal/api"
	"go-openapi-starter/internal/gen"
	"log"
	"net/http"
)

func main() {
	r := chi.NewRouter()

	// Initialize your API handlers
	handlers := &api.Handlers{}

	// Use the generated HandlerFromMux to wire up your handlers
	gen.HandlerFromMux(handlers, r)

	// Serve the OpenAPI spec
	r.Get("/openapi.json", func(w http.ResponseWriter, r *http.Request) {
		http.ServeFile(w, r, "api/openapi.json")
	})

	log.Println("🚀 Server running on :8080")
	log.Fatal(http.ListenAndServe(":8080", r))
}
