package utils

import (
	"strconv"
)

func StrToInt(input string) (int, error) {
	res, err := strconv.Atoi(input)
	return res, err
}

func StrToUInt64(input string) (uint64, error) {
	res, err := strconv.Atoi(input)
	return uint64(res), err
}
