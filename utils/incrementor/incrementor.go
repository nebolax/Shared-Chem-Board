package incrementor

import "ChemBoard/utils/configs"

var values map[string]interface{} = map[string]interface{}{}

func init() {
	if g := configs.Get("incrementor"); g != nil {
		values = g.(map[string]interface{})
	}
}

func Next(key string, save bool) uint64 {
	if values[key] == nil {
		values[key] = 0.0
	}
	values[key] = values[key].(float64) + 1
	configs.Set("incrementor", values)
	return uint64(values[key].(float64))
}
