package deb

import (
	"fmt"
	"path/filepath"
	"sync"

	"github.com/baojingh/prctl/internal/grpool"
	"github.com/baojingh/prctl/internal/utils/files"
)

type DebComponentMeta struct {
	Distribution string
	Component    string
	Architech    string
}

// input: /xx/xx/xx/, check is it exists
func UploadDeb(meta DebComponentMeta, input string) {
	// arch := meta.Architech
	// dis := meta.Distribution
	// com := meta.Component
	var wg sync.WaitGroup
	fileList, _ := files.ListFilesInDir(input)
	for _, file := range fileList {
		wg.Add(1)
		fileName := file
		f := func() {
			doUpload(input, fileName)
			wg.Done()
		}
		grpool.SubmitTask(f)
	}

}

func doUpload(path string, name string) {
	absFilePath := filepath.Join(path, name)
	fmt.Print(absFilePath)
}
