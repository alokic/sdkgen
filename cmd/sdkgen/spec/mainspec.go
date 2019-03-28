package spec

import (
	"encoding/json"
	"io/ioutil"
	"os"

	"github.com/alokic/sdkgen/openapi"
	"github.com/pkg/errors"
)

// ReadMainSpec reads main.json
func ReadMainSpec(path string) (*openapi.MainSpec, error) {
	f, err := os.Open(path)
	if err != nil {
		return nil, errors.Wrap(err, "main.json failed to open")
	}

	js, _ := ioutil.ReadAll(f)

	m := &openapi.MainSpec{}
	json.Unmarshal(js, &m)

	return m, nil
}
