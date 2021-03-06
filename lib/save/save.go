package save

import (
	"encoding/json"
	"fmt"
	"io"
	"io/ioutil"
	"net/http"
	"os"
	"regexp"

	"github.com/bitly/go-simplejson"
	"github.com/webcyou-org/coppeno/coppeno"
	"github.com/webcyou-org/coppeno/lib/load"
	"github.com/webcyou-org/coppeno/utils"
)

func Start(group string, name string, targetPath string) error {
	coppenoJson := load.Start()
	coppeno := coppeno.Coppeno{
		Name: name,
		Url:  targetPath,
	}

	if len(group) > 0 {
		coppenoJson.Set(group, append(coppenoJson.Get(group).MustArray(), coppeno))
	} else {
		coppenoJson.Set("none", append(coppenoJson.Get("none").MustArray(), coppeno))
	}

	f, err := os.Create(utils.CoppenoJsonPath())
	if err != nil {
		return err
	}
	defer f.Close()

	jsonEncode := json.NewEncoder(f)
	jsonEncode.SetIndent("", "    ")
	err = jsonEncode.Encode(coppenoJson)
	if err != nil {
		return err
	}
	return nil
}

func File(filePath string) error {
	r := regexp.MustCompile("^https?://")
	if r.MatchString(filePath) {
		return SaveHostingFile(filePath)
	}
	return SaveLocalFile(filePath)
}

func SaveLocalFile(filePath string) error {
	raw, err := ioutil.ReadFile(filePath)
	if err != nil {
		fmt.Println(err.Error())
		os.Exit(1)
	}

	js, err := simplejson.NewJson(raw)
	if err != nil {
		return err
	}
	return SaveJson(js)
}

func SaveHostingFile(url string) error {
	downloadUrl := utils.GetDownloadUrl(url)
	resp, err := http.Get(downloadUrl)
	if err != nil {
		return err
	}
	defer resp.Body.Close()
	body, _ := io.ReadAll(resp.Body)

	js, err := simplejson.NewJson(body)
	if err != nil {
		return err
	}
	return SaveJson(js)
}

func SaveJson(js *simplejson.Json) error {
	if len(js.MustMap()) > 0 {
		if len(js.Get("name").MustString()) > 0 {
			Start("", js.Get("name").MustString(), js.Get("url").MustString())
		} else {
			for key, _ := range js.MustMap() {
				for i, _ := range js.Get(key).MustArray() {
					Start(key, js.Get(key).GetIndex(i).Get("name").MustString(), js.Get(key).GetIndex(i).Get("url").MustString())
				}
			}
		}
	} else if len(js.MustArray()) > 0 {
		for i, _ := range js.MustArray() {
			Start("", js.GetIndex(i).Get("name").MustString(), js.GetIndex(i).Get("url").MustString())
		}
	}
	return nil
}
