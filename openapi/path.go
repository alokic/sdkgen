package openapi

import (
	"encoding/json"
	"fmt"
	"strings"

	"regexp"

	"github.com/alokic/sdkgen/pkg/stringutils"
	"github.com/jinzhu/inflection"
)

var (
	pathExp = regexp.MustCompile("\\{(.*?)\\}")
)

func (p *Path) build() {
	v := strings.ToLower(p.spec.HTTPMethod)

	switch v {
	case "post", "put", "patch", "delete":
		p.ApiDefs[v] = p.buildPost()
	case "get":
		p.ApiDefs[v] = p.buildGet()
	}
}

func (p *Path) buildPost() *apiDef {
	return &apiDef{
		Desciption:  p.spec.Name,
		Tags:        []string{fmt.Sprintf("%v", p.spec.Version)},
		OperationID: p.spec.Operation, // fmt.Sprintf("%s_%s", p.spec.Operation, inflection.Singular(p.spec.HTTPResource)),
		Parameters:  p.createParameters(),
		Request:     p.createRequest(),
		Responses:   p.createResponse(),
	}
}

func (p *Path) buildGet() *apiDef {
	return &apiDef{
		Desciption:  p.spec.Name,
		Tags:        []string{fmt.Sprintf("%v", p.spec.Version)},
		OperationID: p.spec.Operation, // fmt.Sprintf("%s_%s", p.spec.Operation, inflection.Singular(p.spec.HTTPResource)),
		Parameters:  p.createParameters(),
		Responses:   p.createResponse(),
	}
}

func (p *Path) createResponse() map[string]*response {
	key := fmt.Sprintf("%s%s", inflection.Singular(stringutils.ToCamelCase(p.spec.Operation)), inflection.Singular(stringutils.ToCamelCase(p.spec.HTTPResource)))
	return map[string]*response{
		"200": &response{
			Desciption: "Success response",
			Content:    createContent(fmt.Sprintf("#/components/schemas/%sResponse", key), p.spec.Response),
		},
		"4XX": &response{
			Desciption: "Error in processing request",
			Content:    createContent(fmt.Sprintf("#/components/schemas/Error"), p.spec.Error),
		},
		"5XX": &response{
			Desciption: "Server error in processing request",
			Content:    createContent(fmt.Sprintf("#/components/schemas/Error"), p.spec.Error),
		},
	}
}

func (p *Path) createRequest() *request {
	key := fmt.Sprintf("%s%s", inflection.Singular(stringutils.ToCamelCase(p.spec.Operation)), inflection.Singular(stringutils.ToCamelCase(p.spec.HTTPResource)))
	return &request{
		Content:  createContent(fmt.Sprintf("#/components/schemas/%sRequest", key), p.spec.Request),
		Required: true,
	}
}

func (p *Path) createParammeter(key string, typ string, paramType string) *parameter {
	var required bool

	switch paramType {
	case "header":
		required = isIncluded(p.spec.RequiredHeaders, key)
	case "query":
		required = isIncluded(p.spec.RequiredQuery, key)
	case "path":
		required = true
	}

	return &parameter{
		Name:     key,
		In:       paramType,
		Required: required,
		Schema: map[string]interface{}{
			"type": typ,
		},
		Description: key,
	}
}
func (p *Path) createParameters() []*parameter {
	ps := []*parameter{}
	for k, v := range p.spec.Headers {
		typ, _ := typeOf(v)
		ps = append(ps, p.createParammeter(k, typ, "header"))
	}

	for k, v := range p.spec.Query {
		typ, _ := typeOf(v)
		ps = append(ps, p.createParammeter(k, typ, "query"))
	}

	for _, v := range p.extractPathParameters() {
		ps = append(ps, p.createParammeter(v, "string", "path"))
	}

	return ps
}

func (p *Path) extractPathParameters() []string {
	match := pathExp.FindAllStringSubmatch(p.spec.Path, -1)

	ms := []string{}
	for _, m := range match {
		ms = append(ms, m[1])
	}
	return ms
}

func createContent(key string, examples ...interface{}) map[string]interface{} {
	a := &appJSON{
		Schema: map[string]string{"$ref": key},
	}
	if len(examples) == 0 {
		return map[string]interface{}{"application/json": a}
	}

	b, _ := json.Marshal(examples[0])
	a.Example = string(b)

	return map[string]interface{}{"application/json": a}
}
