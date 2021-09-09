package load

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"
)

type Coppeno struct {
	Name string `json:"name"`
	Url  string `json:"url"`
}

func Start(fileName string, fileGroup string) []Coppeno {
	fmt.Println("load: Start")

	raw, err := ioutil.ReadFile("coppeno.json")
	if err != nil {
		fmt.Println(err.Error())
		os.Exit(1)
	}

	var coppeno []Coppeno
	json.Unmarshal(raw, &coppeno)

	return coppeno
}
