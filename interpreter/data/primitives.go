// Internal representation of data

package data

const (
	INTEGER_TYPE = "INTEGER"
	BOOLEAN_TYPE = "BOOLEAN"
	NULL_TYPE    = "NULL"
)

type DataType string

type Data interface {
	Type() DataType
	Inspect() string
}
