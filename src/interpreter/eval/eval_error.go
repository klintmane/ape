package eval

import (
	"fmt"

	"github.com/ape-lang/ape/src/data"
)

func evalError(str string, rest ...interface{}) *data.Error {
	return &data.Error{Message: fmt.Sprintf(str, rest...)}
}
