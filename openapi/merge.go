package openapi

// merge openapis
func merge(oapis []*OAPI) *OAPI {
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
	path, method, apidef := urlMethodAPIDef(source.Paths)
	m, ok := dest.Paths[path].(map[string]interface{})
	if !ok {
		dest.Paths[path] = map[string]interface{}{method: apidef}
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
	for path, apid := range paths {
		pm := apid.(map[string]*apiDef)
		for httpmethod, apidef := range pm {
			return path, httpmethod, apidef
		}
	}
	return "", "", nil
}
