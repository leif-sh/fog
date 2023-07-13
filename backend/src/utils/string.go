package utils

import (
	"strconv"
)

func StrToInt(input string) (int, error) {
	res, err := strconv.Atoi(input)
	return res, err
}
