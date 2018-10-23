// Internal representation of data

package data

const (
	INTEGER_TYPE = "INTEGER"
	BOOLEAN_TYPE = "BOOLEAN"
	NULL_TYPE    = "NULL"
	RETURN_TYPE  = "RETURN_TYPE"
	ERROR_TYPE   = "ERROR_TYPE"
)

type DataType string

type Data interface {
	Type() DataType
	Inspect() string
}
