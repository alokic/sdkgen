package openapi

// Spec of an API
type Spec struct {
	Name            string                 `json:"name"`
	Version         string                 `json:"version"`
	HTTPMethod      string                 `json:"http_method"`
	HTTPResource    string                 `json:"http_resource"`
	Operation       string                 `json:"operation"`
	Path            string                 `json:"path"`
	Headers         map[string]interface{} `json:"headers"`
	Query           map[string]interface{} `json:"query"`
	Request         interface{}            `json:"request"`
	Response        interface{}            `json:"response"`
	Error           map[string]interface{} `json:"error"`
	RequiredRequest []string               `json:"required_request"`
	RequiredQuery   []string               `json:"required_query"`
	RequiredHeaders []string               `json:"required_headers"`
	Meta            map[string]interface{} `json:"meta"`
}

// ToOpenAPI converts an Spec spec to OpenAPI.
func ToOpenAPI(specs ...*Spec) (*OAPI, error) {
	os := []*OAPI{}

	for _, s := range specs {
		o, err := create(s)
		if err != nil {
			return nil, err
		}
		os = append(os, o)
	}

	return merge(os), nil
}
