package copy

import (
	"io"
	"os"
	"path"
)

func Start(targetPath string) error {
	exePath, _ := os.Getwd()
	targetFile := path.Join(exePath, targetPath)
	stockPath := path.Join(exePath, "stock")

	f, err := os.Open(targetFile)
	if err != nil {
		return err
	}
	defer f.Close()

	info, err := f.Stat()
	if err != nil {
		return err
	}

	distPath := path.Join(stockPath, info.Name())
	dst, err := os.Create(distPath)
	if err != nil {
		return err
	}
	defer dst.Close()

	if _, err := io.Copy(dst, f); err != nil {
		return err
	}
	return nil
}