package main

import (
	"log"
	"net/http"

	"github.com/go-chi/chi"
	"www.github.com/shnartho/shipment-handler-go/pkg/handler"
)

func main() {
	r := chi.NewRouter()

	r.Get("/", handler.IndexHandler)
	r.Post("/check", handler.SubmitHandler)

	fs := http.FileServer(http.Dir("../static/"))
	r.Handle("/static/*", http.StripPrefix("/static/", fs))

	// Start the HTTP server on port 8087
	addr := ":8087"
	log.Printf("Server starting on %s...\n", addr)
	log.Fatal(http.ListenAndServe(addr, r))
}
