package services

import (
	"encoding/json"
	"fmt"
	"net/http"

	t "backend/mytypes"
	r "backend/repositories"
)

// GetCodeByID returns a code entity by its ID (must be a valid Dgraph UID)
func GetCodeByID(uid string) t.Code {
	code := r.GetCodeByID(uid)
	return code
}

// GetAllCode returns all code entities as an array of t.Code
func GetAllCode() []t.Code {
	codes := r.GetAllCode()
	return codes
}

// GetCodeFromRequest gets the code and decodes it properly for later usage
func GetCodeFromRequest(r *http.Request) t.Code {
	var code t.Code
	decoder := json.NewDecoder(r.Body)

	err := decoder.Decode(&code)
	if err != nil {
		panic(err)
	}

	return code
}

// SaveCodeInDatabase saves the code as a node in the database
func SaveCodeInDatabase(savedCode t.Code) string {
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
