package openapi

import (
	"fmt"
	"strings"

	"github.com/alokic/sdkgen/pkg/typeutils"
)

// MainSpec struct
type MainSpec struct {
	OpenAPIVersion string                   `json:"openapi" yaml:"openapi"`
	Tags           []map[string]interface{} `json:"tags" yaml:"tags"`
	Info           *info                    `json:"info" yaml:"info"`
	Servers        []map[string]interface{} `json:"servers" yaml:"servers"`
	Paths          map[string]interface{}   `json:"paths" yaml:"paths"`
	oapiInfo       []*OAPIInfo
	spec           map[string]interface{}
}

// FilePathHandler returns file path of an openapi
type FilePathHandler func() string

// OAPIInfo holds information about openapi
type OAPIInfo struct {
	Handler FilePathHandler
	OAPI    *OAPI
}

// NewMainSpec is contructor
func NewMainSpec(oapiInfo []*OAPIInfo, spec map[string]interface{}) *MainSpec {
	return &MainSpec{
		oapiInfo: oapiInfo,
		spec:     spec,
	}
}

// Build mainspec
func (m *MainSpec) Build() {
	inf := m.spec["info"].(map[string]interface{})

	m.OpenAPIVersion = typeutils.ToStr(m.spec["openapi"])

	m.Info = &info{
		Title:       typeutils.ToStr(inf["title"]),
		Description: typeutils.ToStr(inf["description"]),
		Version:     fmt.Sprintf("v%v", typeutils.ToStr(inf["version"])),
		Contact:     inf["contact"].(map[string]interface{}),
	}
	m.Tags = m.toMapSlice("tags")

	m.Servers = m.toMapSlice("servers")
	m.buildPaths()

}

func (m *MainSpec) buildPaths() {
	pathFolder := make(map[string]interface{})
	for _, o := range m.oapiInfo {
		for _, u := range o.OAPI.GetPathURLs() {
			pathFolder[u] = map[string]string{"$ref": o.Handler() + "#" + "/paths/" + strings.ReplaceAll(u, "/", "~1")}
		}
	}

	refs := map[string]interface{}{}
	for url, ref := range pathFolder {
		refs[url] = ref
	}
	m.Paths = refs
}

func (m *MainSpec) toMapSlice(key string) []map[string]interface{} {
	var ss []map[string]interface{}
	for _, v := range m.spec[key].([]interface{}) {
		ss = append(ss, v.(map[string]interface{}))
	}
	return ss
}

func mainSpecPath(folder string) string {
	cs := strings.Split(folder, "/")
	return fmt.Sprintf("./%s/%s.yaml", cs[len(cs)-2], cs[len(cs)-1])
}
