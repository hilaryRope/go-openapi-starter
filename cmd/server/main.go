package main

import (
	"go-openapi-starter/internal/api"
	"log"
	"net/http"
)

func main() {
	r := api.NewRouter()

	log.Println("🚀 Server running on :8080")
	if err := http.ListenAndServe(":8080", r); err != nil {
		log.Fatal(err)
	}
}
