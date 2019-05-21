package openapi

import (
	"github.com/alokic/sdkgen/pkg/stringutils"
)

// NormalizeName normalize plural name..
// not necessary singular
func NormalizeName(str string) string {
	return stringutils.ToCamelCase(str)
}
