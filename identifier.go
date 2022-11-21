package typescript

import "github.com/schema-cafe/go-types/ast"

func Identifier(id ast.Identifier) string {
	return id.Name
}
