package deb

import (
	"fmt"
	"net/http"
	"os"
	"path/filepath"
	"sync"

	"github.com/baojingh/prctl/pkg/files"
	"github.com/baojingh/prctl/pkg/grpool"
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

func (cli *Client) doUpload(meta DebComponentMeta, path string, fileName string) {
	log.Infof("meta: %v, path: %s, fileName: %s", meta, path, fileName)
	arch := meta.Architech
	dis := meta.Distribution
	com := meta.Component
	uploadUrl := fmt.Sprintf("%s/%s;deb.distribution=%s;deb.component=%s;deb.architecture=%s",
		cli.RepoUrl, fileName, dis, com, arch)

	filePath := filepath.Join(path, fileName)
	file, err := os.Open(filePath)
	if err != nil {
		log.Error(err)
		return
	}
	defer file.Close()

	req, err := http.NewRequest("PUT", uploadUrl, file)
	if err != nil {
		log.Fatal(err)
	}
	req.Header.Set("Content-Type", "application/octet-stream")
	req.SetBasicAuth(cli.Username, cli.Password)

	client := GetHttpClient()
	resp, err := client.Do(req)
	if err != nil {
		log.Fatal(err)
	}
	defer resp.Body.Close()

	// // curl -u${USER}:${TOKEN} \
	// //      -XPUT  \
	// //     "${URL}/${file_name};deb.distribution=${DISTRIBUTION};deb.component=${COMPONENT};deb.architecture=${ARCH}" \
	// //     -T "${file}"

	// 打印响应
	log.Infof("Response Status: %s\n", resp.Status)
}
