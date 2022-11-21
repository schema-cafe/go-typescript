package typescript

import "github.com/schema-cafe/go-types/ast"

func Type(t ast.Type) string {
	if t.IsArray {
		return Identifier(t.BaseType) + "[]"
	}
	if t.IsMap {
		return "{[key: string]: " + Identifier(t.BaseType) + "}"
	}
	return Identifier(t.BaseType)
}
