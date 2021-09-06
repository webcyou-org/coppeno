package update

import (
	"encoding/json"
	"fmt"
	"os"
)

type Coppeno struct {
	Name string `json:"name"`
	Url string  `json:"url"`
}

func Start(fileName string, fileGroup string) error {
	f, err := os.Open("coppeno.json")
	if err != nil {
		return err
	}
	defer f.Close()

	var coppeno []*Coppeno
	err = json.NewDecoder(f).Decode(&coppeno)
	if err != nil {
		return err
	}

	for _, v := range coppeno {
		fmt.Printf("%+v\n", v)
	}
	// fmt.Printf("%+v\n", coppeno)
	fmt.Println(fileName)
	fmt.Println(fileGroup)
	return nil
}
