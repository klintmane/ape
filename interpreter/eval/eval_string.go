package eval

import "ape/interpreter/data"

func evalString(value string) *data.String {
	return &data.String{Value: value}
}
