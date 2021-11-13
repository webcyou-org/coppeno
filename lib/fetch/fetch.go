package fetch

import (
	"fmt"
	"io"
	"net/http"
	"os"
	"regexp"
	"strings"

	"github.com/webcyou-org/coppeno/lib/load"
)

type Coppeno struct {
	Name string `json:"name"`
	Url  string `json:"url"`
}

func Start(fileName string, fileGroup string) error {
	coppenoJson := load.Start()

	for key, _ := range coppenoJson.MustMap() {
		for i, _ := range coppenoJson.Get(key).MustArray() {
			url := coppenoJson.Get(key).GetIndex(i).Get("url").MustString()
			name := coppenoJson.Get(key).GetIndex(i).Get("name").MustString()

			// github
			r := regexp.MustCompile("^https?://github.com")
			if r.MatchString(url) {
				downloadUrl := r.ReplaceAllString(url, "https://raw.githubusercontent.com")
				downloadUrl = strings.Replace(downloadUrl, "/blob", "", 1)
				if err := DownloadFile(name, downloadUrl); err != nil {
					panic(err)
				}
				fmt.Println(downloadUrl)
			}

			// bitbucket
			// https://bitbucket.org/<username>/<repo-name>/raw/<branch>/<file-name>
			r = regexp.MustCompile("^https?://bitbucket.org")
			if r.MatchString(url) {
				downloadUrl := strings.Replace(url, "/src", "/raw", 1)
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
