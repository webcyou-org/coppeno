package list

import (
	"fmt"
	"webcyou-org/coppeno/lib/load"
)

func Start() error {
	coppenoJson := load.Start()

	for _, c := range coppenoJson {
		str := fmt.Sprintf("name: %s, url: %s", c.Name, c.Url)
		fmt.Println(str)
	}
	return nil
}
