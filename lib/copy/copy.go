package copy

import (
	"fmt"
	"github.com/markbates/pkger"
	"io"
	"os"
	"path"
)

func Start(targetPath string) error {
	exePath, _ := os.Getwd()
	targetFile := path.Join(exePath, targetPath)

	fmt.Println("targetPath: ", targetPath)
	fmt.Println("targetFile: ", targetFile)

	f, err := pkger.Open(targetFile)
	if err != nil {
		return err
	}

	defer f.Close()

	info, err := f.Stat()
	if err != nil {
		return err
	}

	fmt.Println("Name: ", info.Name())
	fmt.Println("Size: ", info.Size())
	fmt.Println("Mode: ", info.Mode())
	fmt.Println("ModTime: ", info.ModTime())

	if _, err := io.Copy(os.Stdout, f); err != nil {
		return err
	}
	return nil
}