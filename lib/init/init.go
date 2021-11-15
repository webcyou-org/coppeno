package init

import (
	"encoding/json"
	"os"

	"github.com/webcyou-org/coppeno/utils"
)

func Start() error {
	return InitCoppenoJson()
}

func InitCoppenoJson() error {
	var initJson string = `
		{
			"none": []
		}
		`

	f, err := os.Create(utils.CoppenoJsonPath())
	if err != nil {
		return err
	}
	defer f.Close()

	jsonEncode := json.NewEncoder(f)
	jsonEncode.SetIndent("", "    ")
	err = jsonEncode.Encode(initJson)
	if err != nil {
		return err
	}
	return nil
}
