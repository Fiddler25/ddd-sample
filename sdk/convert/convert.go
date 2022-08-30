package convert

import (
	"golang.org/x/exp/constraints"
	"strconv"
)

func StrToInt[T constraints.Integer](val string) T {
	integer, _ := strconv.Atoi(val)
	return T(integer)
}
