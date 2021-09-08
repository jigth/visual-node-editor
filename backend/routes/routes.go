package routes

import (
	c "../controllers"
	"github.com/go-chi/chi"
)

// AddRoutes sets every routes to be used in the API
func AddRoutes(r *chi.Mux) {
	r.Get("/", c.Index)
	r.Get("/code/{codeID:[0-9]+}", c.GetCode)
	r.Get("/code/execute/{codeID:[0-9]+}", c.ExecuteCode)
	r.Post("/code/execute/directly", c.ExecuteCodeDirectly)
	r.Post("/code/save", c.SaveCode)
}
