package openapi

import (
	"fmt"
	"reflect"
	"strings"
)

// isIncluded returns if array includes a key.
func isIncluded(arr []string, key string) bool {
	for _, v := range arr {
		if key == v {
			return true
		}
	}
	return false
}

// typeOf returns type of a value.
func typeOf(value interface{}) (string, error) {
	o := strings.ToLower(reflect.TypeOf(value).Kind().String())

	switch {
	case strings.Contains(o, "int"):
		return "integer", nil
	case strings.Contains(o, "float"):
		return "number", nil
	case strings.Contains(o, "string"):
		return "string", nil
	case strings.Contains(o, "bool"):
		return "boolean", nil
	case strings.Contains(o, "array"), strings.Contains(o, "slice"):
		return "array", nil
	case strings.Contains(o, "map"):
		return "object", nil
	default:
		return "", fmt.Errorf("invalid type for value: %v, type: %v", value, o)
	}
}
