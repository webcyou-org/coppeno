package load

import (
	"fmt"
	"io/ioutil"
	"os"

	"github.com/bitly/go-simplejson"
	"github.com/webcyou-org/coppeno/utils"
)

func Start() *simplejson.Json {
	if _, err := os.Stat(utils.CoppenoJsonPath()); err != nil {
		var coppenoJson string = `
		{
			"none": []
		}
		`
		jsonBytes := []byte(coppenoJson)

		os.Mkdir(utils.CoppenoPath(), 0775)
		err := ioutil.WriteFile(utils.CoppenoJsonPath(), jsonBytes, 0775)
		if err != nil {
			fmt.Println(err.Error())
			os.Exit(1)
		}
	}

	raw, err := ioutil.ReadFile(utils.CoppenoJsonPath())
	if err != nil {
		fmt.Println(err.Error())
		os.Exit(1)
	}

	json, _ := simplejson.NewJson(raw)
	return json
}
