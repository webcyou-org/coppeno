package load

import (
	"encoding/json"
	"fmt"
	"io"
	"io/ioutil"
	"os"
	"path"
)

type Coppeno struct {
	Name string `json:"name"`
	Url  string `json:"url"`
}

func Start() []Coppeno {
	fmt.Println("load file - coppeno.json")

	raw, err := ioutil.ReadFile("coppeno.json")
	if err != nil {
		fmt.Println(err.Error())
		os.Exit(1)
	}

	var coppeno []Coppeno
	json.Unmarshal(raw, &coppeno)

	return coppeno
}

func File(filePath string) error {
	file, err := os.Open(filePath)
	if err != nil {
		return err
	}
	defer file.Close()

	exePath, _ := os.Getwd()
	targetFile := path.Join(exePath, "coppeno.json")

	dst, err := os.Create(targetFile)
	if err != nil {
		return err
	}
	defer dst.Close()

	if _, err := io.Copy(dst, file); err != nil {
		return err
	}
	return nil
}
