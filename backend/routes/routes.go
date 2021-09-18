package routes

import (
	c "backend/controllers"

	"github.com/go-chi/chi"
)

// AddRoutes sets every routes to be used in the API
func AddRoutes(r *chi.Mux) {
	r.Get("/", c.Index)
	r.Get("/code", c.GetAllCode)
	r.Get("/code/", c.GetAllCode)
	r.Get("/code/{codeID:[0-9x]+}", c.GetCodeByID)
	r.Get("/code/execute/{codeID:[0-9x]+}", c.ExecuteCode)
	r.Post("/code/execute/directly", c.ExecuteCodeDirectly)
	r.Post("/code/save", c.SaveCode)
}
