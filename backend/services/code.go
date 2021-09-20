package services

import (
	"encoding/json"
	"net/http"

	t "backend/mytypes"
	repo "backend/repositories"
)

// GetCodeByID returns a code entity by its ID (must be a valid Dgraph UID)
func GetCodeByID(uid string) t.Code {
	code := repo.GetCodeByID(uid)
	return code
}

// GetAllCode returns all code entities as an array of t.Code
func GetAllCode() []t.Code {
	codes := repo.GetAllCode()
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

// SaveCodeInDatabase saves (or updates) the code as a node in the database
func SaveCodeInDatabase(newCode t.Code) string {
	savedCode := repo.SaveCode(newCode)

	res, _ := json.Marshal(savedCode)
	return string(res)
}
