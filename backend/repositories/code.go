package repositories

import (
	db "backend/database"
	t "backend/mytypes"
	"context"
	"encoding/json"
	"fmt"
	"log"

	"github.com/dgraph-io/dgo/v210"
)

var client *dgo.Dgraph = db.Connect()

// GetAllCode returns all code entities as an array of t.Code
func GetAllCode() []t.Code {
	txn := client.NewTxn()
	ctx := context.Background()
	defer txn.Discard(ctx)
	query := `
		{
			 data( func: has(Code.name)){
				 CodeID: uid
				 Name: Code.name
				 Code: Code.code
				 AstTree: Code.astTree
			 }
		}
	`
	resp, err := txn.Query(context.Background(), query)
	if err != nil {
		log.Fatal(err)
	}

	var codes struct {
		Data []t.Code
	}

	if err := json.Unmarshal(resp.GetJson(), &codes); err != nil {
		log.Fatal(err)
	}
	return codes.Data
}

// GetCodeByID returns a code entity by its ID (must be a valid Dgraph UID)
func GetCodeByID(uid string) t.Code {
	txn := client.NewTxn()
	ctx := context.Background()
	defer txn.Discard(ctx)
	q := `
		{
			 data( func: uid(%s)){
				 CodeID: uid
				 Name: Code.name
				 Code: Code.code
				 AstTree: Code.astTree
			 }
		}
	`

	query := fmt.Sprintf(q, uid)
	resp, err := txn.Query(context.Background(), query)
	if err != nil {
		log.Fatal(err)
	}

	var codes struct {
		Data []t.Code
	}

	if err := json.Unmarshal(resp.GetJson(), &codes); err != nil {
		log.Fatal(err)
	}

	return codes.Data[0]
}
