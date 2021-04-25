package configs

import (
	"encoding/json"
	"io/fs"
	"io/ioutil"
	"os"
)

var (
	data map[string]interface{} = map[string]interface{}{}
)

func init() {
	os.OpenFile("configs.txt", os.O_RDONLY|os.O_CREATE, os.ModeDevice)
	inp, err := ioutil.ReadFile("configs.txt")
	if err != nil {
		panic(err.Error())
	} else {
		err := json.Unmarshal(inp, &data)
		if err != nil {
			println(err.Error())
		}
	}
}

func Get(key string) interface{} {
	return data[key]
}

func Set(key string, value interface{}) {
	data[key] = value
	b, err := json.Marshal(data)
	if err != nil {
		println(err.Error())
	} else {
		err := ioutil.WriteFile("configs.txt", b, fs.ModeDevice)
		if err != nil {
			println(err)
		}
	}
}
