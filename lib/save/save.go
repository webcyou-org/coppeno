package save

import (
	"encoding/json"
	"fmt"
	"os"
	"webcyou-org/coppeno/lib/load"
)

type Coppeno struct {
	Name string `json:"name"`
	Url  string `json:"url"`
}

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
	coppenoList := append(coppenoJson, coppeno)

	f, err := os.Create("data_output.json")
	if err != nil {
		return err
	}
	defer f.Close()

	err = json.NewEncoder(f).Encode(coppenoList)
	if err != nil {
		return err
	}
	return nil
}
