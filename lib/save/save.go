package save

import (
	"encoding/json"
	"fmt"
	"os"
	"webcyou-org/coppeno/lib/load"
)

func Start(name string, targetPath string) error {
	fmt.Println("save: Start")

	coppenoJson := load.Start("fileName", "fileGroup")
	// for _, c := range coppenoJson {
	// 	// fmt.Sprintf("%s: %s", c.Name, c.Url)
	// 	fmt.Println(c.Name)
	// 	fmt.Println(c.Url)
	// }
	coppeno := load.Coppeno{
		Name: name,
		Url:  targetPath,
	}
	coppenoJson = append(coppenoJson, coppeno)

	f, err := os.Create("coppeno.json")
	if err != nil {
		return err
	}
	defer f.Close()

	err = json.NewEncoder(f).Encode(coppenoJson)
	if err != nil {
		return err
	}
	return nil
}
