package utils

import "github.com/valyala/fastjson"

func JsonDumps(input string) *fastjson.Value {
	value, err := fastjson.Parse(input)
	if err != nil {
		SugarLogger.Error("error json", "json", input)
		return nil
	}
	return value
}
