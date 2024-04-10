package files

import (
	"fmt"
	"testing"
)

func TestCreateDirIfNotExist(t *testing.T) {
	path := "./aa/bb/c.txt"
	err := CreateDirIfNotExist(path, 0755)
	if err != nil {
		fmt.Printf("Failed to create path, %s", err)
	}
}

func TestCreateFileIfNotExist(t *testing.T) {
	path := "a.txt"
	err := CreateFileIfNotExist(path, 0600)
	if err != nil {
		fmt.Printf("Failed to create file, %s", err)
	}
}
