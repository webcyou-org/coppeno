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
	stockPath := path.Join(exePath, "stock")
	fmt.Println(stockPath)

	f, err := pkger.Open(targetFile)
	if err != nil {
		return err
	}

	defer f.Close()

	info, err := f.Stat()
	if err != nil {
		return err
	}

	exe, err := os.Executable()
	fmt.Println("exe: ", exe)
	fmt.Println("Name: ", info.Name())
	fmt.Println("Size: ", info.Size())
	fmt.Println("Mode: ", info.Mode())
	fmt.Println("ModTime: ", info.ModTime())

	distPath := path.Join(stockPath, info.Name())
	dst, err := os.Create(distPath)
	if err != nil {
		panic(err)
	}
	defer dst.Close()

	if _, err := io.Copy(dst, f); err != nil {
		return err
	}
	return nil
}