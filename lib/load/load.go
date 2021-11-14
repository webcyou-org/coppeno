package load

import (
	"fmt"
	"io/ioutil"
	"os"

	"github.com/bitly/go-simplejson"
)

func Start() *simplejson.Json {
	raw, err := ioutil.ReadFile("coppeno.json")
	if err != nil {
		fmt.Println(err.Error())
		os.Exit(1)
	}

	json, _ := simplejson.NewJson(raw)
	return json
}
