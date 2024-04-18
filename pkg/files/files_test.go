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

func TestListFilesInDir(t *testing.T) {
	res, _ := ListFilesInDir("./")
	fmt.Println(res)
}

func TestComposeAbsPath(t *testing.T) {
	list := []string{"c", "d", "e"}

	res := ComposeAbsPath("/hello/world/", list)
	fmt.Println(res)
}

func TestMoveFilesBatch(t *testing.T) {
	src := "/data/code/goproject/prctl/examples/src-deb"
	dst := "/data/code/goproject/prctl/examples/dst-deb"
	MoveFilesBatch(src, dst, ".deb")
}

func TestRemoveFileIfExist(t *testing.T) {
	res := RemoveFileIfExist("/home/secur1ty/.prctl/ee")
	fmt.Println(res)
}

func TestGetFileNameFromAbsPath(t *testing.T) {
	name := GetFileNameFromAbsPath("/data/code/goproject/prctl/pkg/files/files_test.go")
	fmt.Println(name)
}
