package list

import (
	"fmt"
	"webcyou-org/coppeno/lib/load"
)

type Coppeno struct {
	Name string `json:"name"`
	Url  string `json:"url"`
}

func Start() error {
	coppenoJson := load.Start("fileName", "fileGroup")

	for _, c := range coppenoJson {
		str := fmt.Sprintf("name: %s, url: %s", c.Name, c.Url)
		fmt.Println(str)
	}
	return nil
}
