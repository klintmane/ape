package eval

import "github.com/ape-lang/ape/src/data"

// Global references, so a new object does not get allocated for each evaluation
var (
	TRUE  = &data.Boolean{Value: true}
	FALSE = &data.Boolean{Value: false}
)

func evalBoolean(value bool) *data.Boolean {
	if value {
		return TRUE
	}
	return FALSE
}
