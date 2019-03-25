package openapi

import (
	"fmt"

	"github.com/alokic/sdkgen/pkg/stringutils"
	"github.com/alokic/sdkgen/pkg/typeutils"
	"github.com/jinzhu/inflection"
)

// Builder buiils an openapi spec.
type Builder struct {
	Spec   *Spec
	schema *Schema
	path   *Path
	oapi   *OAPI
}

// NewBuilder is constructor.
func NewBuilder(s *Spec) *Builder {
	return &Builder{
		schema: newSchema(s),
		path:   newPath(s),
		Spec:   s,
		oapi:   &OAPI{},
	}
}

// Meta builds meta info
func (b *Builder) Meta() (*Builder, error) {
	inf := b.Spec.Meta["info"].(map[string]interface{})

	b.oapi.OpenAPIVersion = typeutils.ToStr(b.Spec.Meta["openapi"])
	b.oapi.Info = &info{
		Title:       fmt.Sprintf(typeutils.ToStr(inf["title"]), inflection.Singular(stringutils.ToCamelCase(b.Spec.HTTPResource))),
		Description: fmt.Sprintf(typeutils.ToStr(inf["description"]), inflection.Singular(stringutils.ToCamelCase(b.Spec.HTTPResource))),
		Version:     b.version(),
		Contact:     inf["contact"].(map[string]interface{}),
	}
	b.oapi.Tags = []map[string]interface{}{
		map[string]interface{}{"name": b.Spec.HTTPResource},
	}
	b.oapi.Servers = b.toServer()

	return b, nil
}

// Component builds component.
func (b *Builder) Component() (*Builder, error) {
	err := b.schema.build("Request", b.Spec.Request, fmt.Sprintf("%sRequest", scope(b.Spec.Operation, b.Spec.HTTPResource)))
	if err != nil {
		return nil, err
	}

	err = b.schema.build("Response", b.Spec.Response, fmt.Sprintf("%sResponse", scope(b.Spec.Operation, b.Spec.HTTPResource)))
	if err != nil {
		return nil, err
	}

	err = b.schema.build("Error", b.Spec.Error, "Error")
	if err != nil {
		return nil, err
	}

	b.oapi.Components = &component{Schemas: b.schema.root}

	return b, nil
}

// Path builds component.
func (b *Builder) Path() (*Builder, error) {
	b.path.build()
	b.oapi.Paths = map[string]interface{}{b.Spec.URL: b.path.ApiDefs}
	return b, nil
}

// Build openpi spec
func (b *Builder) Build() *OAPI {
	return b.oapi
}

func (b *Builder) toServer() []map[string]interface{} {
	var ss []map[string]interface{}
	for _, v := range b.Spec.Meta["servers"].([]interface{}) {
		ss = append(ss, v.(map[string]interface{}))
	}
	return ss
}

func (b *Builder) version() string {
	return fmt.Sprintf("%v", b.Spec.Version)
}

func scope(op, resource string) string {
	return stringutils.ToCamelCase(op) + inflection.Singular(stringutils.ToCamelCase(resource))
}
