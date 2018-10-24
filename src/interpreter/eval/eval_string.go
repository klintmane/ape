package eval

import "ape/src/interpreter/data"

func evalString(value string) *data.String {
	return &data.String{Value: value}
}
