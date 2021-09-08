package controllers

import (
	"encoding/json"
	"fmt"
	"net/http"

	"../services"

	"github.com/go-chi/chi"
)

// Index is the Controller for the index route
func Index(w http.ResponseWriter, r *http.Request) {
	res, _ := json.Marshal([]string{"Welcome", "Adventurer"})
	fmt.Fprint(w, `{"Message":`+string(res)+`}`)
}

// GetCode gets the python code from the database
func GetCode(w http.ResponseWriter, r *http.Request) {
	codeID := chi.URLParam(r, "codeID")

	// Response structure
	type getResponse struct {
		Status string
		Code   string `json:"code"`
	}

	// Response instance
	newResponse := &getResponse{
		Status: "Code with id " + codeID + " executed successfully",
		Code:   "Code with id " + codeID + " goes here",
	}

	// Serialize response
	res, _ := json.Marshal(newResponse)

	fmt.Fprintf(w, string(res))
}

// ExecuteCode calls a python interpreter and sends the response after executing
// the code
func ExecuteCode(w http.ResponseWriter, r *http.Request) {
	codeID := chi.URLParam(r, "CodeID")
	code := chi.URLParam(r, "Code")
	fmt.Println(code)
	fmt.Fprintf(w, string(`{"status": "Code with id %v executed successfully", "result": "14"}`), codeID)
}

// ExecuteCode calls a python interpreter and sends the response after executing
// the code
func ExecuteCodeDirectly(w http.ResponseWriter, r *http.Request) {
	code := services.GetCodeFromRequest(r).Code
	fmt.Println(code)
	result := services.ExecutePythonCode(code)
	fmt.Fprintf(w, result)
}

// SaveCode saves the python code into the database
func SaveCode(w http.ResponseWriter, r *http.Request) {
	code := services.GetCodeFromRequest(r)
	res := services.SaveCodeInDatabase(code)
	fmt.Fprintf(w, res)
}
