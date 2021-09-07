package save

import (
	"encoding/json"
	"fmt"
	"os"
	"webcyou-org/coppeno/lib/load"
)

func Start(name string, targetPath string) error {
	coppeno := load.Start("fileName", "fileGroup")

	fmt.Println(coppeno)
	fmt.Println(name)
	fmt.Println(targetPath)
	a := map[string]interface{}{
		"name": name,
		"url": targetPath,
	}

	f, err := os.Create("data_output.json")
	if err != nil {
		return err
	}
	defer f.Close()

	err = json.NewEncoder(f).Encode(a)
	if err != nil {
		return err
	}
	return nil
}