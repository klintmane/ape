package eval

import "github.com/ape-lang/ape/src/data"

func evalBoolean(value bool) *data.Boolean {
	if value {
		return data.TRUE
	}
	return data.FALSE
}
