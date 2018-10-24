package data

import "hash/fnv"

type String struct {
	Value string
}

func (s *String) Type() DataType { return STRING_TYPE }

func (s *String) Inspect() string { return s.Value }

func (s *String) Hash() uint64 {
	h := fnv.New64a()
	h.Write([]byte(s.Value))
	return h.Sum64()
}
