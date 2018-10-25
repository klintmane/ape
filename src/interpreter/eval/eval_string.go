package eval

import "ape/src/data"

func evalString(value string) *data.String {
	return &data.String{Value: value}
}
