package main

import (
	"log"
	"net/http"

	"github.com/go-chi/chi/v5"

	"greeter/internal/gen"
	"greeter/internal/handlers"
)

func main() {
	r := chi.NewRouter()
	srv := handlers.NewServer()
	handler := gen.HandlerWithOptions(gen.NewStrictHandler(srv, nil), gen.ChiServerOptions{BaseRouter: r})

	log.Println("greeter listening on :9090")
	if err := http.ListenAndServe(":9090", handler); err != nil {
		log.Fatalf("server stopped: %v", err)
	}
}
