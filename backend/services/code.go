package services

import (
	"encoding/json"
	"fmt"
	"net/http"

	t "../types"
)

func GetCodeFromRequest(r *http.Request) t.Code {
	var code t.Code
	decoder := json.NewDecoder(r.Body)

	err := decoder.Decode(&code)
	if err != nil {
		panic(err)
	}

	return code
}

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
