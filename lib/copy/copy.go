package copy

import (
	"fmt"
	"os"
	"path"
	"io"
	"github.com/markbates/pkger"
)

func Start(dirName string, targetPath string, assetsType string) error {
	exePath, _ := os.Getwd()
	createPath := path.Join(exePath, targetPath)

	fmt.Println("dirName: ", dirName)
	fmt.Println("dirName: ", dirName)
	fmt.Println("createPath: ", createPath)

	f, err := pkger.Open(createPath)
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