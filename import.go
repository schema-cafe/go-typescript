package typescript

import "github.com/schema-cafe/go-types/ast"

func Import(id ast.Identifier) string {
	return "import " + id.Name + " from '" + id.Path + "';"
}
