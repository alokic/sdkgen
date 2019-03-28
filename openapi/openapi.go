package openapi

// OAPI struct
type OAPI struct {
	OpenAPIVersion string                   `json:"openapi" yaml:"openapi"`
	Tags           []map[string]interface{} `json:"tags" yaml:"tags"`
	Info           *info                    `json:"info" yaml:"info"`
	Servers        []map[string]interface{} `json:"servers" yaml:"servers"`
	Paths          map[string]interface{}   `json:"paths" yaml:"paths"`
	Components     *component               `json:"components" yaml:"components"`
}

type info struct {
	Title       string                 `json:"title" yaml:"title"`
	Description string                 `json:"description" yaml:"description"`
	Version     string                 `json:"version" yaml:"version"`
	Contact     map[string]interface{} `json:"contact" yaml:"contact"`
}

// Path structure
type Path struct {
	ApiDefs map[string]*apiDef
	spec    *Spec
}

type apiDef struct {
	Desciption  string               `json:"description" yaml:"description"`
	Tags        []string             `json:"tags" yaml:"tags"`
	Parameters  []*parameter         `json:"parameters" yaml:"parameters"`
	OperationID string               `json:"operationId" yaml:"operationId"`
	Request     *request             `json:"requestBody,omitempty" yaml:"requestBody,omitempty"`
	Responses   map[string]*response `json:"responses" yaml:"responses"`
}

type parameter struct {
	Name        string                 `json:"name" yaml:"name"`
	In          string                 `json:"in" yaml:"in"`
	Required    bool                   `json:"required" yaml:"required"`
	Schema      map[string]interface{} `json:"schema" yaml:"schema"`
	Description string                 `json:"description" yaml:"description"`
}

type request struct {
	Content map[string]interface{} `json:"content" yaml:"content"`
}

type response struct {
	Desciption string                 `json:"description" yaml:"description"`
	Content    map[string]interface{} `json:"content" yaml:"content"`
}

type appJSON struct {
	Schema  map[string]string `json:"schema" yaml:"schema"`
	Example string            `json:"example,omitempty" yaml:"example,omitempty"`
}

type component struct {
	Schemas map[string]*Node `json:"schemas" yaml:"schemas"`
}

// getPathURLs returns urls in path
func (o *OAPI) getPathURLs() []string {
	var urls []string
	for url := range o.Paths {
		urls = append(urls, url)
	}
	return urls
}

func newComponent() *component {
	return &component{
		Schemas: make(map[string]*Node),
	}
}

func newPath(s *Spec) *Path {
	return &Path{spec: s, ApiDefs: make(map[string]*apiDef)}
}

// toOpenAPI converts an Spec spec to OpenAPI.
func create(s *Spec) (*OAPI, error) {
	b := NewBuilder(s)
	_, err := b.Meta()
	if err != nil {
		return nil, err
	}

	_, err = b.Path()
	if err != nil {
		return nil, err
	}

	_, err = b.Component()
	if err != nil {
		return nil, err
	}

	return b.Build(), nil
}
