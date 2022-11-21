package typescript

import (
	"strings"

	"github.com/schema-cafe/go-types/form"
)

func Interface(name string, fields []form.Field) string {
	s := strings.Builder{}
	s.WriteString("interface ")
	s.WriteString(name)
	s.WriteString(" {\n")
	for _, f := range fields {
		s.WriteString("  ")
		s.WriteString(f.Name)
		s.WriteString(": ")
		s.WriteString(Type(f.Type))
		s.WriteString(";\n")
	}
	s.WriteString("}\n")
	return s.String()
}
