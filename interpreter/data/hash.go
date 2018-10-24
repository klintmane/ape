package data

import (
	"fmt"
	"strings"
)

type HashKey struct {
	Type  DataType
	Value uint64
}

type HashPair struct {
	Key   Data
	Value Data
}

type Hash struct {
	Pairs map[HashKey]HashPair
}

func (b *Hash) Type() DataType { return HASH_TYPE }

func (h *Hash) Inspect() string {
	var sb strings.Builder
	pairs := []string{}

	for _, pair := range h.Pairs {
		pairs = append(pairs, fmt.Sprintf("%s: %s",
			pair.Key.Inspect(), pair.Value.Inspect()))
	}

	sb.WriteString("{")
	sb.WriteString(strings.Join(pairs, ", "))
	sb.WriteString("}")
	return sb.String()
}

// Interface
type HashableData interface {
	Data
	Hash() uint64
}

func HashData(h HashableData) HashKey {
	return HashKey{Type: h.Type(), Value: h.Hash()}
}
