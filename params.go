package golwf

type Param struct {
	Key   string `json: key`
	Value string `json: value`
}

type Params []Param

func (params Params) Get(key string) (string, bool) {
	for _, param := range params {
		if key == param.Key {
			return param.Value, true
		}
	}

	return "", false
}

func (params *Params) Set(key string, value string) {
	*params = append(*params, Param{key, value})
}

func (params *Params) Reset() {
	*params = (*params)[:0]
}
