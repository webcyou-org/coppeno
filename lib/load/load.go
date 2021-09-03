package load

import (
	"encoding/json"
	"fmt"
	"os"
)

var coppeno struct {
	Name string `json: "name"`
	Url string  `json: "url"`
}

func Start(fileName string, fileGroup string) error {
	f, err := os.Open("coppeno.json")
	if err != nil {
		return err
	}
	defer f.Close()

	err = json.NewDecoder(f).Decode(&coppeno)
	if err != nil {
		return err
	}
	fmt.Printf("%+v\n", coppeno)
	fmt.Println(fileName)
	fmt.Println(fileGroup)
	return nil
}
