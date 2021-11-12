package save

import (
	"encoding/json"
	"fmt"
	"os"
	"webcyou-org/coppeno/lib/load"
)

func Start(group string, name string, targetPath string) error {
	fmt.Println("save: Start")

	coppenoJson := load.Start()
	coppeno := load.Coppeno{
		Name: name,
		Url:  targetPath,
	}

	if len(group) > 0 {
		coppenoJson.Set(group, append(coppenoJson.Get(group).MustArray(), coppeno))
	} else {
		coppenoJson.Set("none", append(coppenoJson.Get("none").MustArray(), coppeno))
	}

	f, err := os.Create("coppeno.json")
	if err != nil {
		return err
	}
	defer f.Close()

	jsonEncode := json.NewEncoder(f)
	jsonEncode.SetIndent("", "    ")
	err = jsonEncode.Encode(coppenoJson)
	if err != nil {
		return err
	}
	return nil
}
