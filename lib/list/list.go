package list

import (
	"fmt"
	"webcyou-org/coppeno/lib/load"
)

func Start() error {
	coppenoJson := load.Start()

	for key, _ := range coppenoJson.MustMap() {
		fmt.Println(fmt.Sprintf("group: %s", key))
		for i, _ := range coppenoJson.Get(key).MustArray() {
			str := fmt.Sprintf(
				"  name: %s   url: %s",
				coppenoJson.Get(key).GetIndex(i).Get("name").MustString(),
				coppenoJson.Get(key).GetIndex(i).Get("url").MustString())
			fmt.Println(str)
		}
	}
	return nil
}
