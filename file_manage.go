package helpers

import (
	"os"
	"path"
)

func WriteFile(dir string, name string, data []byte) error {
	dirErr := os.MkdirAll(dir, 0755)
	if dirErr != nil {
		return dirErr
	}
	path := path.Join(dir, name)
	file, err := os.Create(path)
	if err != nil {
		return err
	}
	defer file.Close()
	_, err = file.Write(data)
	if err != nil {
		return err
	}
	return nil
}

func DeleteFile(dir string, name string) error {
	path := path.Join(dir, name)
	err := os.Remove(path)
	if err != nil {
		return err
	}
	return nil
}
