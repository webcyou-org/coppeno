package save

import (
	"encoding/json"
	"fmt"
	"os"
)

func Start(key string, targetPath string) error {
	fmt.Println(key)
	fmt.Println(targetPath)
	a := map[string]interface{}{
		"x": 123,
		"y": "ABC",
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