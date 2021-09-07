package controllers

import (
	"encoding/json"
	"fmt"
	"net/http"

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
	codeID := chi.URLParam(r, "codeID")
	fmt.Fprintf(w, string(`{"status": "Code with id %v executed successfully", "result": "14"}`), codeID)
}

type code struct {
	CodeID string
	Name   string
	Code   string
}

// SaveCode saves the python code into the database
func SaveCode(w http.ResponseWriter, r *http.Request) {
	var code code
	decoder := json.NewDecoder(r.Body)

	err := decoder.Decode(&code)
	if err != nil {
		panic(err)
	}

	res := saveInDatabase(code)
	fmt.Fprintf(w, res)
}

func saveInDatabase(savedCode code) string {
	code := savedCode.Code
	name := savedCode.Name
	id := savedCode.CodeID
	fmt.Println()
	fmt.Println("Saving...")
	fmt.Println("code: ", code)
	fmt.Println("name: ", name)
	fmt.Println("id: ", id)
	fmt.Println("Saved to the database succesfully")
	fmt.Println()

	res, _ := json.Marshal(savedCode)
	return string(res)
}
