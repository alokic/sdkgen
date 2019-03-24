package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"
	"path"
	"strings"

	"github.com/alokic/sdkgen/openapi"
	"github.com/pkg/errors"
	"gopkg.in/yaml.v2"
)

var (
	outputFolder = "./xyz"
)

func writeYaml(v interface{}, p string) error {
	d, err := yaml.Marshal(v)
	if err != nil {
		return err
	}

	err = os.MkdirAll(path.Dir(p), os.ModePerm)
	if err != nil {
		return err
	}

	return ioutil.WriteFile(p, d, 0644)
}

func errlog(err error) {
	if err == nil {
		return
	}
	fmt.Println(err)
	os.Exit(1)
}

func openAPIPath(folder string) string {
	cs := strings.Split(folder, "/")
	return fmt.Sprintf("%s/%s/%s.yaml", outputFolder, cs[len(cs)-2], cs[len(cs)-1])
}

func mainSpecPath(folder string) string {
	cs := strings.Split(folder, "/")
	return fmt.Sprintf("./%s/%s.yaml", cs[len(cs)-2], cs[len(cs)-1])
}

// toOpenAPI converts an Spec spec to OpenAPI.
func toOpenAPI(s *openapi.Spec) (*openapi.OAPI, error) {
	b := openapi.NewBuilder(s)
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

// toMainSpec converts an Spec spec to main.yaml.
func toMainSpec(oapiInfo []*openapi.OAPIInfo) (*openapi.MainSpec, error) {
	path := "spec/json" + "/main.json"
	f, err := os.Open(path)
	if err != nil {
		return nil, errors.Wrap(err, "main.json failed to open")
	}

	js, _ := ioutil.ReadAll(f)

	var m map[string]interface{}
	json.Unmarshal(js, &m)

	ms := openapi.NewMainSpec(oapiInfo, m)
	ms.Build()

	return ms, nil
}

func main() {
	specs, err := openapi.ReadSpecs("spec/json")

	oapiInfo := []*openapi.OAPIInfo{}
	for folder, ss := range specs {
		os := []*openapi.OAPI{}

		for _, s := range ss {
			o, err := toOpenAPI(s)
			errlog(err)
			os = append(os, o)
		}

		o := openapi.Merge(os)
		errlog(writeYaml(o, openAPIPath(folder)))

		oapiInfo = append(oapiInfo, &openapi.OAPIInfo{
			OAPI: o,
			Handler: func(folder string) openapi.FilePathHandler {
				return func() string {
					return mainSpecPath(folder)
				}
			}(folder),
		})
	}

	m, err := toMainSpec(oapiInfo)
	errlog(err)
	writeYaml(m, outputFolder+"/main.yaml")
}
