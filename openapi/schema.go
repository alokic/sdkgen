package openapi

import (
	"fmt"
	"reflect"
	"strings"

	"github.com/alokic/sdkgen/pkg/stringutils"
	"github.com/pkg/errors"
)

// Schema struct.
type Schema struct {
	root map[string]*Node
	spec *Spec
}

var gschema *Schema

// Node struct
type Node struct {
	Properties           map[string]*Node `json:"properties,omitempty" yaml:"properties,omitempty"`
	Items                *Node            `json:"items" yaml:"items,omitempty"`
	AdditionalProperties bool             `json:"additionalProperties,omitempty" yaml:"additionalProperties,omitempty"`
	Type                 string           `json:"type" yaml:"type,omitempty"`
	Ref                  string           `json:"ref" yaml:"$ref,omitempty"`
	Required             []string         `json:"required" yaml:"required,omitempty"`
	context              string
}

// newSchema is constructor.
func newSchema(s *Spec) *Schema {
	gschema = &Schema{
		root: map[string]*Node{},
		spec: s,
	}
	return gschema
}

// build schema of key, val
func (s *Schema) build(key string, value interface{}, ctx string) error {
	typ, err := typeOf(value)
	if err != nil {
		return errors.Wrapf(err, "Generate")
	}
	n := newNode(ctx)
	err = n.generate(ctx, value, typ)

	if key == "Request" {
		s.addRequiredParams(ctx)
	}
	return err
}

func (s *Schema) addRequiredParams(ctx string) {
	for _, v := range s.spec.RequiredRequest {
		os := strings.Split(v, ".")

		newctx := ctx
		if len(os) > 1 {
			parent := os[len(os)-2] // most recent parent of the key(os[len(os)-1])
			newctx = objectName(ctx, parent)
		}

		key := os[len(os)-1]
		s.root[newctx].Required = append(s.root[newctx].Required, key)
	}
}

// Get schema
func (s *Schema) get() map[string]*Node {
	return s.root
}

func newNode(ctx string) *Node {
	return &Node{
		Properties: make(map[string]*Node),
		context:    ctx,
	}
}

func (n *Node) add(key string, in *Node) {
	if n.Type == "" {
		if in.Type == "object" {
			in.AdditionalProperties = true
		}

		gschema.root[key] = in
	}

	if n.Type == "object" {
		n.Properties[key] = in
	}

	if n.Type == "array" {
		n.Items = in
	}
}

// generate openapi schema
func (n *Node) generate(key string, value interface{}, valueType string) error {
	switch valueType {
	case "object":
		newn := newNode(n.context)
		newn.Type = "object"

		err := newn.properties(value)
		if err != nil {
			return errors.Wrapf(err, "generate/object: key=%v, value=%v", key, value)
		}

		n.add(key, newn)

	case "array":
		newn := newNode(n.context)
		newn.Type = "array"

		err := newn.array(fmt.Sprintf("%sItem", strings.Title(key)), arrayElement(value))
		if err != nil {
			return errors.Wrapf(err, "generate/array: key=%v, value=%v", key, value)
		}

		n.add(key, newn)

	case "nestedObject":
		newn := newNode(n.context)
		newn.Ref = fmt.Sprintf("#/components/schemas/%s", objectName(n.context, stringutils.ToCamelCase(key)))

		n.add(key, newn)

		root := newNode(n.context)
		root.generate(objectName(n.context, key), value, "object")
	default:
		n.add(key, &Node{Type: valueType})
	}
	return nil
}

func (n *Node) properties(value interface{}) error {
	for k, v := range value.(map[string]interface{}) {
		err := n.processType(k, v)
		if err != nil {
			return errors.Wrapf(err, "properties: key=%v, value=%v", k, v)
		}
	}
	return nil
}

func (n *Node) array(key string, value interface{}) error {
	err := n.processType(key, value)
	if err != nil {
		return errors.Wrapf(err, "array: value=%v", value)
	}
	return nil
}

func (n *Node) processType(key string, value interface{}) error {
	typ, err := typeOf(value)
	if err != nil {
		return errors.Wrap(err, "processType/typeOf")
	}

	if typ == "object" {
		typ = "nestedObject"
	}

	err = n.generate(key, value, typ)
	if err != nil {
		return errors.Wrap(err, "processType/generate")
	}

	return nil
}

func objectName(ctx, key string) string {
	return ctx + stringutils.ToCamelCase(key)
}

func arrayElement(value interface{}) interface{} {
	s := reflect.ValueOf(value)
	for i := 0; i < s.Len(); i++ {
		return s.Index(i).Interface()
	}
	return "" // return string in case of nil or empty array
}
