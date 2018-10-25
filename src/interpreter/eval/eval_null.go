package eval

import "ape/src/data"

// Global references, so a new object does not get allocated for each evaluation
var (
	NULL = &data.Null{}
)

func evalNull() *data.Null {
	return NULL
}
