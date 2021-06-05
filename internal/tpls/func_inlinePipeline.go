package tpls

func init() {
	funcs["inlinePipeline"] = func(keyAndValues ...interface{}) map[string]interface{} {
		var result = make(map[string]interface{})

		if len(keyAndValues)%2 != 0 {
			panic("tpls/inlinePipeline: must have an even number of arguments")
		}

		for i := 0; i < len(keyAndValues); i += 2 {
			key, ok := keyAndValues[i].(string)
			if !ok {
				panic("tpls/inlinePipeline: key must be string")
			}
			result[key] = keyAndValues[i+1]
		}

		return result
	}
}
