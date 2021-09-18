package main

import (
	"backend/routes"
	"net/http"

	"github.com/go-chi/chi"
)

func main() {
	r := chi.NewRouter()
	routes.ConfigureRouter(r)
	routes.AddRoutes(r)

	http.ListenAndServe(":3000", r)
}
