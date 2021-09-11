package save

import (
	"encoding/json"
	"fmt"
	"io"
	"os"
	"path"
	"webcyou-org/coppeno/lib/load"
)

func Start(name string, targetPath string) error {
	fmt.Println("save: Start")

	coppenoJson := load.Start("fileName", "fileGroup")
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
