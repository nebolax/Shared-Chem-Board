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
	os.OpenFile("configs.conf", os.O_RDONLY|os.O_CREATE, os.ModeDevice)
	inp, err := ioutil.ReadFile("configs.conf")
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
		err := ioutil.WriteFile("configs.conf", b, fs.ModeDevice)
		if err != nil {
			println(err)
		}
	}
}
