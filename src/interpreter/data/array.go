package data

import "strings"

type Array struct {
	Elements []Data
}

func (ao *Array) Type() DataType { return ARRAY_TYPE }

func (ao *Array) Inspect() string {
	var sb strings.Builder
	elements := []string{}

	for _, e := range ao.Elements {
		elements = append(elements, e.Inspect())
	}

	sb.WriteString("[")
	sb.WriteString(strings.Join(elements, ", "))
	sb.WriteString("]")

	return sb.String()
}
