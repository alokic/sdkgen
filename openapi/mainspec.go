package openapi

import (
	"strings"
)

// MainSpec struct
type MainSpec struct {
	OpenAPIVersion string                   `json:"openapi" yaml:"openapi"`
	Tags           []map[string]interface{} `json:"tags" yaml:"tags"`
	Info           *info                    `json:"info" yaml:"info"`
	Servers        []map[string]interface{} `json:"servers" yaml:"servers"`
	Paths          map[string]interface{}   `json:"paths" yaml:"paths"`
}

// Build mainspec
func (m *MainSpec) Build(filePath string, oapi *OAPI) {
	if m.Paths == nil {
		m.Paths = map[string]interface{}{}
	}

	for _, u := range oapi.getPathURLs() {
		m.Paths[u] = map[string]string{"$ref": filePath + "#" + "/paths/" + strings.ReplaceAll(u, "/", "~1")}
	}
}
