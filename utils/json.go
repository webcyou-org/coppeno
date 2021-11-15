package utils

import (
	"io/ioutil"
	"os"
)

func CreateCoppenoJson() error {
	var coppenoJson string = `
		{
			"none": []
		}
		`
	jsonBytes := []byte(coppenoJson)

	os.Mkdir(CoppenoPath(), 0775)
	err := ioutil.WriteFile(CoppenoJsonPath(), jsonBytes, 0775)
	if err != nil {
		return err
	}
	return nil
}
