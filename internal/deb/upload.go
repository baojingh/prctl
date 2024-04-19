package deb

import (
	"fmt"
	"net/http"
	"os"
	"path/filepath"
	"sync"

	"github.com/baojingh/prctl/pkg/files"
	"github.com/baojingh/prctl/pkg/grpool"
	"github.com/baojingh/prctl/pkg/prhttp"
)

type DebComponentMeta struct {
	Distribution string
	Component    string
	Architech    string
}

func (cli *Client) Upload(meta DebComponentMeta, input string) {
	log.Infof("start upload, input path %s", input)
	var wg sync.WaitGroup

	fileList, _ := files.ListFilesInDir(input)
	for _, file := range fileList {
		fileName := file
		wg.Add(1)
		f := func() {
			defer wg.Done()
			cli.doUpload(meta, input, fileName)
		}
		grpool.SubmitTask(f)
	}
	// NOTE: Do Not Forget it.
	wg.Wait()
}

//	curl -u${USER}:${TOKEN} \
//	     -XPUT  \
//	    "${URL}/${file_name};deb.distribution=${DISTRIBUTION};deb.component=${COMPONENT};deb.architecture=${ARCH}" \
//	    -T "${file}"
func (cli *Client) doUpload(meta DebComponentMeta, path string, fileName string) {
	arch := meta.Architech
	dis := meta.Distribution
	com := meta.Component
	uploadUrl := fmt.Sprintf("%s/%s;deb.distribution=%s;deb.component=%s;deb.architecture=%s",
		cli.RepoUrl, fileName, dis, com, arch)

	// Open the file
	filePath := filepath.Join(path, fileName)
	file, err := os.Open(filePath)
	if err != nil {
		log.Error(err)
		return
	}
	defer file.Close()

	// Create request object
	req, err := http.NewRequest("PUT", uploadUrl, file)
	if err != nil {
		log.Error(err)
		return
	}

	// Set username and password
	req.Header.Set("Content-Type", "application/octet-stream")
	req.SetBasicAuth(cli.Username, cli.Password)

	// Do request
	resp, err := prhttp.DoHttpRequest(req)
	if err != nil {
		log.Error(err)
		return
	}
	defer resp.Body.Close()

	log.Infof("Response Status: %s\n", resp.Status)
}
