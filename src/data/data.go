// Internal representation of data

package data

const (
	INTEGER_TYPE           = "INTEGER"
	BOOLEAN_TYPE           = "BOOLEAN"
	NULL_TYPE              = "NULL"
	RETURN_TYPE            = "RETURN"
	ERROR_TYPE             = "ERROR"
	FUNCTION_TYPE          = "FUNCTION"
	STRING_TYPE            = "STRING"
	BUILTIN_TYPE           = "BUILTIN"
	ARRAY_TYPE             = "ARRAY"
	HASH_TYPE              = "HASH"
	COMPILED_FUNCTION_TYPE = "COMPILED_TYPE"
	CLOSURE_TYPE           = "CLOSURE_TYPE"
)

// Global references, so a new object does not get allocated for each evaluation
var (
	TRUE  = &Boolean{Value: true}
	FALSE = &Boolean{Value: false}
	NULL  = &Null{}
)

type DataType string

type Data interface {
	Type() DataType
	Inspect() string
}
