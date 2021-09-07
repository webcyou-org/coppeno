package save

import (
	"encoding/json"
	"fmt"
	"os"
)

func Start(name string, targetPath string) error {
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