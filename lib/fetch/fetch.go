package fetch

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"os"
	"regexp"
	"strings"
)

type Coppeno struct {
	Name string `json:"name"`
	Url  string `json:"url"`
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
		// fmt.Printf("%+v\n", v)
		fmt.Println(v.Name)
		fmt.Println(v.Url)

		r := regexp.MustCompile("^https?://github.com")
		if r.MatchString(v.Url) {
			downloadUrl := r.ReplaceAllString(v.Url, "https://raw.githubusercontent.com")
			downloadUrl = strings.Replace(downloadUrl, "/blob", "", 1)
			if err := DownloadFile(v.Name, downloadUrl); err != nil {
				panic(err)
			}
			fmt.Println(downloadUrl)
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
