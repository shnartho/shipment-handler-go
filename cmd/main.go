package main

import (
	"log"
	"net/http"
	"os"
	"path/filepath"

	"github.com/go-chi/chi"
	"www.github.com/shnartho/shipment-handler-go/pkg/handler"
)

func main() {
	exePath, err := os.Executable()
	if err != nil {
		log.Fatalf("Error: %v", err)
		return
	}
	exeDir := filepath.Dir(exePath)

	// Set the working directory to the executable's directory
	err = os.Chdir(exeDir)
	if err != nil {
		log.Fatalf("Error: %v", err)
		return
	}

	r := chi.NewRouter()

	r.Get("/", handler.IndexHandler)
	r.Get("/order", handler.OrderHandler)
	r.Get("/add", handler.AddHandler)
	r.Get("/update", handler.UpdateHandler)
	r.Get("/remove", handler.RemoveHandler)

	r.Post("/orderpack", handler.OrderPackHandler)
	r.Post("/addpack", handler.AddPackHandler)
	r.Post("/updatepack", handler.UpdatePackHandler)
	r.Post("/removepack", handler.RemovePackHandler)

	r.Get("/getslicevalues", handler.GetSliceHandler)

	fs := http.FileServer(http.Dir("../static/"))
	r.Handle("/static/*", http.StripPrefix("/static/", fs))

	// Start the HTTP server on port 8087
	addr := ":8087"
	log.Printf("Server starting on %s...\n", addr)
	log.Fatal(http.ListenAndServe(addr, r))
}
