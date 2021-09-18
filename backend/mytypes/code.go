// Use this file to put all your types or structs related to the "code" domain

package mytypes

// Code is an entity used to store the data (incluiding the AST) of the code in the DB
type Code struct {
	CodeID  string
	Name    string
	Code    string
	AstTree string
}
