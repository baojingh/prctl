package files

import (
	"errors"
	"os"
)

func CreateDirIfNotExist(path string, perm os.FileMode) error {
	err := os.MkdirAll(path, os.ModePerm)
	if err != nil {
		return err
	}
	return nil
}

// O_RDONLY : Open a file for read only operations
// O_WRONLY : Open a file for write only operations
// O_RDWR : Open a file for read-write
// O_APPEND :It appends data to the file when writing
// O_CREATE: It creates a file if none exists.
func CreateFileIfNotExist(path string, perm os.FileMode) error {
	file, err := os.OpenFile(path, os.O_RDONLY|os.O_CREATE, perm)
	if err != nil {
		return err
	}
	return file.Close()
}

func IsFileExist(path string) bool {
	_, err := os.Stat(path)
	if err == nil {
		return true
	}
	if errors.Is(err, os.ErrNotExist) {
		return false
	}
	return false
}

func RemoveFileIfExist(path string) bool {
	if IsFileExist(path) {
		os.Remove(path)
		return true
	}
	return false
}
