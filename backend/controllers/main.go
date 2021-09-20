package controllers

import (
	"backend/services"
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/go-chi/chi"
)

const codeID = "codeID"

// Index is the Controller for the index route
func Index(w http.ResponseWriter, r *http.Request) {
	res, _ := json.Marshal([]string{"Welcome", "Adventurer"})
	fmt.Fprint(w, `{"Message":`+string(res)+`}`)
}

// GetCodeByID gets the python code from the database by its UID
func GetCodeByID(w http.ResponseWriter, r *http.Request) {
	codeID := chi.URLParam(r, codeID)

	resObject := services.GetCodeByID(codeID)
	res, _ := json.Marshal(resObject)

	fmt.Fprintf(w, string(res))
}

// GetAllCode gets all python code nodes (entities) from the database
func GetAllCode(w http.ResponseWriter, r *http.Request) {
	resObject := services.GetAllCode()
	res, _ := json.Marshal(resObject)
	fmt.Fprintf(w, string(res))
}

// ExecuteCode calls a python interpreter and sends the response after executing
// the code. The code is retrieved from the database
func ExecuteCode(w http.ResponseWriter, r *http.Request) {
	codeID := chi.URLParam(r, codeID)

	// Get code from database
	code := services.GetCodeByID(codeID).Code

	// Execute it and return response
	res := services.ExecutePythonCode(string(code))
	fmt.Fprintf(w, res)
}

// ExecuteCodeDirectly executes python code directly passed to the server
// NOTE/WARNING: This function is for testing purposes only, not for production usage
func ExecuteCodeDirectly(w http.ResponseWriter, r *http.Request) {
	// Get code from request
	code := services.GetCodeFromRequest(r).Code

	// Execute it and return response
	result := services.ExecutePythonCode(code)
	fmt.Fprintf(w, result)
}

// SaveCode saves (or updates) the python code into the database
func SaveCode(w http.ResponseWriter, r *http.Request) {
	code := services.GetCodeFromRequest(r)
	res := services.SaveCodeInDatabase(code)
	fmt.Fprintf(w, res)
}
