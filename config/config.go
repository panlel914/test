package config

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
)

type config struct {

}

const(
	Dev = 1
	Test = 2
	Prod = 3
)

type info struct {
	configType string
}

func Load(fileName string, envName string, envType string) (*config, error){
	f, err := ioutil.ReadFile(fileName)
	if err != nil {
		fmt.Print(err)
	}
	map2 := make(map[string]interface{})
	err = json.Unmarshal(f, &map2)
	if err != nil {
		fmt.Println(err)
	}
	return nil, err
}