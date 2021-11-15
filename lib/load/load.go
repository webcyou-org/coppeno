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
		utils.CreateCoppenoJson()
	}

	raw, err := ioutil.ReadFile(utils.CoppenoJsonPath())
	if err != nil {
		fmt.Println(err.Error())
		os.Exit(1)
	}

	json, _ := simplejson.NewJson(raw)
	return json
}
