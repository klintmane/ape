package eval

import (
	"ape/interpreter/data"
	"fmt"
)

func evalError(str string, rest ...interface{}) *data.Error {
	return &data.Error{Message: fmt.Sprintf(str, rest...)}
}
