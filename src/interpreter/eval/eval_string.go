package eval

import "github.com/ape-lang/ape/src/data"

func evalString(value string) *data.String {
	return &data.String{Value: value}
}
