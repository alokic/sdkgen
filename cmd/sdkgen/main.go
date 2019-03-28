package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"
	"path"
	"strings"

	v2config "github.com/alokic/gopkg/config/v2"
	"github.com/alokic/sdkgen/cmd/sdkgen/spec"
	"github.com/alokic/sdkgen/openapi"
	"github.com/go-kit/kit/log"
	"github.com/pkg/errors"
	"gopkg.in/yaml.v2"
)

type svcConfig struct {
	InputPath  string `json:"in" usage:"input folder having json specs of API" required:"true"`
	OutputPath string `json:"out" usage:"output folder where openapi3 specs for APIs are generated" required:"true" print:"true"`
}

var (
	config      = &svcConfig{}
	svcName     = "sdkgen"
	logger      log.Logger
	projectRoot = fmt.Sprintf("%s/src/github.com/alokic/%s", os.Getenv("GOPATH"), svcName)
)

// Build vars
var (
	// GITCOMMIT means git commit tag
	GITCOMMIT = ""
	// GITBRANCH means git branch for the build
	GITBRANCH = ""
	// VERSION means release version
	VERSION = "0.0.0"
)

func initializeLogger() {
	logger = log.NewLogfmtLogger(log.NewSyncWriter(os.Stderr))
	logger = log.With(logger, "ts", log.DefaultTimestampUTC, "caller", log.DefaultCaller, "app", svcName)
}

func initializeConfig() {
	c := v2config.NewConfig(svcName)
	err := c.Load(config)
	if err != nil {
		logger.Log("config", "initialize", "error", err)
		os.Exit(1)
	}
	c.Print(log.With(logger, "dump", "config"))
}

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

func formatOutPath(folder string) string {
	ins := strings.Split(strings.TrimSuffix(strings.TrimPrefix(config.InputPath, "./"), "/"), "/")
	cs := strings.Split(folder, "/")
	p := strings.Join(cs[len(ins):len(cs)], "/")
	if p == "" {
		return "spec"
	}
	return p
}

func openAPIPath(folder string) string {
	return fmt.Sprintf("%s/%s.yaml", config.OutputPath, formatOutPath(folder))
}

func mainSpecPath(folder string) string {
	return fmt.Sprintf("./%s.yaml", formatOutPath(folder))
}

// toMainSpec converts an Spec spec to main.yaml.
func toMainSpec(oapiInfo []*openapi.OAPIInfo) (*openapi.MainSpec, error) {
	path := config.InputPath + "/main.json"
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
	initializeLogger()
	initializeConfig()

	specs, err := spec.Read(config.InputPath)

	oapiInfo := []*openapi.OAPIInfo{}
	for infolder, ss := range specs {
		o, err := openapi.ToOpenAPI(ss...)
		errlog(err)

		errlog(writeYaml(o, openAPIPath(infolder)))

		oapiInfo = append(oapiInfo, &openapi.OAPIInfo{
			OAPI: o,
			Handler: func(infolder string) openapi.FilePathHandler {
				return func() string {
					return mainSpecPath(infolder)
				}
			}(infolder),
		})
	}

	m, err := toMainSpec(oapiInfo)
	errlog(err)
	writeYaml(m, config.OutputPath+"/main.yaml")
}
