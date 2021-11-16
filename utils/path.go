package utils

import (
	"fmt"
	"os"
	"path/filepath"
	"regexp"
	"strings"
)

func CoppenoPath() string {
	conf, err := os.UserConfigDir()
	if err != nil {
		fmt.Fprintln(os.Stderr, err)
	}
	return filepath.Join(conf, "coppeno")
}

func CoppenoJsonPath() string {
	conf, err := os.UserConfigDir()
	if err != nil {
		fmt.Fprintln(os.Stderr, err)
	}
	return filepath.Join(conf, "coppeno", "coppeno.json")
}

func ConfigJsonPath() string {
	conf, err := os.UserConfigDir()
	if err != nil {
		fmt.Fprintln(os.Stderr, err)
	}
	return filepath.Join(conf, "coppeno", "config.json")
}

func GetDownloadUrl(url string) string {
	// github
	r := regexp.MustCompile("^https?://github.com")
	if r.MatchString(url) {
		downloadUrl := r.ReplaceAllString(url, "https://raw.githubusercontent.com")
		return strings.Replace(downloadUrl, "/blob", "", 1)
	}

	// gitlab
	r = regexp.MustCompile("^https?://gitlab")
	if r.MatchString(url) {
		return strings.Replace(url, "/blob", "/raw", 1)
	}

	// bitbucket
	// https://bitbucket.org/<username>/<repo-name>/raw/<branch>/<file-name>
	r = regexp.MustCompile("^https?://bitbucket.org")
	if r.MatchString(url) {
		return strings.Replace(url, "/src", "/raw", 1)
	}

	// Other Hosting
	return url
}
