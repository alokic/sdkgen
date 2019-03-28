package spec

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"
	"path/filepath"

	dirPath "path"

	"github.com/alokic/sdkgen/openapi"
	"github.com/pkg/errors"
)

type requestArray struct {
	Request []interface{} `json:"request"`
}

type responseArray struct {
	Response []interface{} `json:"response"`
}

// Read reads all api specs(as json) in a folder
func Read(folder string) (map[string][]*openapi.Spec, error) {
	apis := map[string][]*openapi.Spec{}

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
			api := &openapi.Spec{}
			if err := json.Unmarshal(spec, api); err != nil {
				return errors.Wrap(err, fmt.Sprintf("folder-walk file:%s json.Unmarshal failed", path))
			}

			if api.Request == nil {
				readRequestArray(api, spec)
			}

			if api.Response == nil {
				readResponseArray(api, spec)
			}

			err = readMeta(api, folder)
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

func readRequestArray(s *openapi.Spec, data []byte) error {
	req := &requestArray{}
	if err := json.Unmarshal(data, req); err != nil {
		return errors.Wrap(err, fmt.Sprintf("failed to read requestArray %s %s s", s.Version, s.HTTPResource))
	}
	s.Request = req.Request
	return nil
}

func readResponseArray(s *openapi.Spec, data []byte) error {
	resp := &responseArray{}
	if err := json.Unmarshal(data, resp); err != nil {
		return errors.Wrap(err, fmt.Sprintf("failed to read responseArray %s %s s", s.Version, s.HTTPResource))
	}
	s.Response = resp.Response
	return nil
}

func readMeta(s *openapi.Spec, folder string) error {
	var meta map[string]interface{}

	b, err := ioutil.ReadFile(fmt.Sprintf("%s/main.json", folder))
	if err != nil {
		return err
	}
	json.Unmarshal(b, &meta)
	s.Meta = meta
	return nil
}
