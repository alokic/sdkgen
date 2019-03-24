package openapi

// Merge openapis
func Merge(oapis []*OAPI) *OAPI {
	if len(oapis) == 0 {
		return nil
	}

	o := &OAPI{
		OpenAPIVersion: oapis[0].OpenAPIVersion,
		Tags:           oapis[0].Tags,
		Info:           oapis[0].Info,
		Servers:        oapis[0].Servers,
		Components:     newComponent(),
		Paths:          map[string]interface{}{},
	}
	for _, oapi := range oapis {
		o = mergePaths(oapi, o)
		o = mergeSchema(oapi, o)
	}

	return o
}

func mergePaths(source, dest *OAPI) *OAPI {
	url, method, apidef := urlMethodAPIDef(source.Paths)
	m, ok := dest.Paths[url].(map[string]interface{})
	if !ok {
		dest.Paths[url] = map[string]interface{}{method: apidef}
	} else {
		m[method] = apidef
	}
	return dest
}

func mergeSchema(source, dest *OAPI) *OAPI {
	nm := dest.Components.Schemas

	for k, v := range source.Components.Schemas {
		nm[k] = v
	}
	return dest
}

func urlMethodAPIDef(paths map[string]interface{}) (string, string, *apiDef) {
	for url, path := range paths {
		pm := path.(map[string]*apiDef)
		for httpmethod, apidef := range pm {
			return url, httpmethod, apidef
		}
	}
	return "", "", nil
}
