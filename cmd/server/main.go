package main

import (
	"net/http"

	"github.com/SolBaa/receipt-processor/cmd/server/routes"
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
)

func main() {

	r := chi.NewRouter()
	r.Use(middleware.Logger)
	routes.Routes(r)

	http.ListenAndServe(":9000", r)
}
