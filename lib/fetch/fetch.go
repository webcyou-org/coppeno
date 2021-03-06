package fetch

import (
	"fmt"
	"io"
	"net/http"
	"os"

	"github.com/webcyou-org/coppeno/lib/load"
	"github.com/webcyou-org/coppeno/utils"
)

func Start(fileName string, fileGroup string) error {
	fmt.Println(fileName)
	fmt.Println(fileGroup)
	All()

	return nil
}

func All() error {
	coppenoJson := load.Start()

	for key, _ := range coppenoJson.MustMap() {
		for i, _ := range coppenoJson.Get(key).MustArray() {
			url := coppenoJson.Get(key).GetIndex(i).Get("url").MustString()
			name := coppenoJson.Get(key).GetIndex(i).Get("name").MustString()

			downloadUrl := utils.GetDownloadUrl(url)
			if err := DownloadFile(name, downloadUrl); err != nil {
				panic(err)
			}
			fmt.Println(downloadUrl)
		}
	}
	return nil
}

func Group(fileGroup string) error {
	coppenoJson := load.Start()

	for i, _ := range coppenoJson.Get(fileGroup).MustArray() {
		url := coppenoJson.Get(fileGroup).GetIndex(i).Get("url").MustString()
		name := coppenoJson.Get(fileGroup).GetIndex(i).Get("name").MustString()

		downloadUrl := utils.GetDownloadUrl(url)
		if err := DownloadFile(name, downloadUrl); err != nil {
			panic(err)
		}
		fmt.Println(downloadUrl)
	}
	return nil
}

func Single(fileName string) error {
	coppenoJson := load.Start()

	for key, _ := range coppenoJson.MustMap() {
		for i, _ := range coppenoJson.Get(key).MustArray() {
			if fileName == coppenoJson.Get(key).GetIndex(i).Get("name").MustString() {
				url := coppenoJson.Get(key).GetIndex(i).Get("url").MustString()
				name := coppenoJson.Get(key).GetIndex(i).Get("name").MustString()

				downloadUrl := utils.GetDownloadUrl(url)
				if err := DownloadFile(name, downloadUrl); err != nil {
					panic(err)
				}
				fmt.Println(downloadUrl)
			}
		}
	}
	return nil
}

func DownloadFile(filepath string, url string) error {
	resp, err := http.Get(url)
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	out, err := os.Create(filepath)
	if err != nil {
		return err
	}
	defer out.Close()

	_, err = io.Copy(out, resp.Body)
	return err
}
