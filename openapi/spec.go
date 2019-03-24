package openapi

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"
	"path/filepath"

	dirPath "path"

	"github.com/pkg/errors"
)

// Spec of an API
type Spec struct {
	Name            string                 `json:"name"`
	Version         string                 `json:"version"`
	HTTPMethod      string                 `json:"http_method"`
	HTTPResource    string                 `json:"http_resource"`
	Operation       string                 `json:"operation"`
	URL             string                 `json:"url"`
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

type requestArray struct {
	Request []interface{} `json:"request"`
}

type responseArray struct {
	Response []interface{} `json:"response"`
}

func (s *Spec) String() string {
	if s == nil {
		return ""
	}
	b, _ := json.Marshal(s)
	return string(b)
}

// ReadSpecs reads all api specs(as json) in a folder
func ReadSpecs(folder string) (map[string][]*Spec, error) {
	apis := map[string][]*Spec{}

	err := filepath.Walk(folder,
		func(path string, info os.FileInfo, err error) error {
			if err != nil {
				return errors.Wrap(err, fmt.Sprintf("folder-walk file:%s failed", path))
			}

			if info.IsDir() || info.Name() == "meta.json" || info.Name() == "main.json" {
				return nil
			}

			f, err := os.Open(path)
			if err != nil {
				return errors.Wrap(err, fmt.Sprintf("folder-walk file:%s open failed", path))
			}

			spec, _ := ioutil.ReadAll(f)
			api := &Spec{}
			if err := json.Unmarshal(spec, api); err != nil {
				return errors.Wrap(err, fmt.Sprintf("folder-walk file:%s json.Unmarshal failed", path))
			}

			if api.Request == nil {
				api.readRequestArray(spec)
			}

			if api.Response == nil {
				api.readResponseArray(spec)
			}

			err = api.readMeta(folder)
			if err != nil {
				return err
			}

			apis[dirPath.Dir(path)] = append(apis[dirPath.Dir(path)], api)

			return nil
		})
	if err != nil {
		return nil, errors.Wrap(err, "folder-walk failed")
	}

	return apis, nil
}

func (s *Spec) readRequestArray(data []byte) error {
	req := &requestArray{}
	if err := json.Unmarshal(data, req); err != nil {
		return errors.Wrap(err, fmt.Sprintf("failed to read requestArray %s %s s", s.Version, s.HTTPResource))
	}
	s.Request = req.Request
	return nil
}

func (s *Spec) readResponseArray(data []byte) error {
	resp := &responseArray{}
	if err := json.Unmarshal(data, resp); err != nil {
		return errors.Wrap(err, fmt.Sprintf("failed to read responseArray %s %s s", s.Version, s.HTTPResource))
	}
	s.Response = resp.Response
	return nil
}

func (s *Spec) readMeta(folder string) error {
	var meta map[string]interface{}

	b, err := ioutil.ReadFile(fmt.Sprintf("%s/meta.json", folder))
	if err != nil {
		return err
	}
	json.Unmarshal(b, &meta)
	s.Meta = meta
	return nil
}
