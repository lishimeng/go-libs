package file

import (
	"os"
	"path/filepath"
	"time"
)

func BuildPath(root string, sub string) string {

	parent := filepath.ToSlash(root)

	f := parent + `/` + sub
	return filepath.FromSlash(f)
}

func RemovePath(path string) (err error) {
	err = os.RemoveAll(path)
	time.Sleep(2 * time.Second)
	return
}

func CreateFolder(path string) error {
	err := os.MkdirAll(path, os.ModePerm)
	time.Sleep(2 * time.Second)
	return err
}
