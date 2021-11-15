package utils

import (
	"fmt"
	"os"
	"path/filepath"
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
