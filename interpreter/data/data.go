// Internal representation of data

package data

const (
	INTEGER_TYPE  = "INTEGER"
	BOOLEAN_TYPE  = "BOOLEAN"
	NULL_TYPE     = "NULL"
	RETURN_TYPE   = "RETURN"
	ERROR_TYPE    = "ERROR"
	FUNCTION_TYPE = "FUNCTION"
	STRING_TYPE   = "STRING"
	BUILTIN_TYPE  = "BUILTIN"
)

type DataType string

type Data interface {
	Type() DataType
	Inspect() string
}
