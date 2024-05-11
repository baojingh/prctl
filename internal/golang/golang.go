package golang

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"os"
	"path/filepath"
	"strings"
	"sync"

	"github.com/baojingh/prctl/internal/common"
	"github.com/baojingh/prctl/internal/handler"
	"github.com/baojingh/prctl/internal/logger"
	"github.com/baojingh/prctl/pkg/files"
	"github.com/baojingh/prctl/pkg/grpool"
	"github.com/baojingh/prctl/pkg/prhttp"
	"github.com/baojingh/prctl/pkg/shell"
)

type GoRepoManage struct {
	handler.Client
}

var log = logger.New()

func NewGoRepository() handler.RepoManage {
	cli := common.CreateClient()
	return &GoRepoManage{Client: *cli}
}

// curl  -u ee:ee  -X DELETE https://rr.rtf-alm.ee.cloud/artifactory/ee-dev-go-ee/pool
func (cli *GoRepoManage) Delete(param handler.DeleteParam) {
	// Create request object
	url := fmt.Sprintf("%s/artifactory/%s/%s", cli.RepoUrl, cli.RepoName, "pool")
	req, err := http.NewRequest("DELETE", url, nil)
	if err != nil {
		log.Errorf("Create request err, %s", err)
		return
	}
	// Set username and password
	req.Header.Set("Content-Type", "text/plain")
	req.SetBasicAuth(cli.Username, cli.Password)

	// Do request
	resp, err := prhttp.DoHttpRequest(req)
	if err != nil {
		log.Errorf("fail to send req, %s", err)
		return
	}
	defer resp.Body.Close()
	log.Infof("delete success, status: %d", resp.StatusCode)
}

// input: /xx/xx/xx/ss.txt, check is it exists
// output aa/ss/ created  by pip command automatically
func (cli *GoRepoManage) Download(input string, output string) {
	// go mod download
	params := []string{"mod", "download"}
	log.Infof("Command: go %s", strings.Join(params, " "))
	out, err := shell.DoShellCmd("go", params...)
	if err != nil {
		log.Errorf("failed to download Components, err: %v", out)
		return
	}
	log.Info("Download success.")
}

func (cli *GoRepoManage) Upload(meta handler.ComponentMeta, input string) {
	log.Infof("start upload, input path %s", input)
	var wg sync.WaitGroup

	fileList, _ := files.ListFilesInDir(input)
	for _, file := range fileList {
		fileName := file
		wg.Add(1)
		f := func() {
			defer wg.Done()
			cli.doUpload(input, fileName)
		}
		grpool.SubmitTask(f)
	}
	// NOTE: Do Not Forget it.
	wg.Wait()
}

//	curl -u${USER}:${TOKEN} \
//	     -XPUT  \
//	    "${URL}/flask/3.0.3/flafmvsni_vnsi7823_u9r32.whl" \
//	    -T "${file}"
func (cli *GoRepoManage) doUpload(path string, fileName string) {

	name := strings.Split(fileName, "-")[0]
	version := strings.Split(fileName, "-")[1]

	uploadUrl := fmt.Sprintf("%s/artifactory/%s/%s/%s/%s",
		cli.RepoUrl, cli.RepoName, name, version, fileName)

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

	log.Infof("HTTP upload Success Status")
}

func (cli *GoRepoManage) List() []handler.ComponentView {
	// Create request object
	url := fmt.Sprintf("%s/%s", cli.RepoUrl, "artifactory/api/search/aql")

	reqB := fmt.Sprintf(`items.find({"repo":"%s", "name":{"$match":"*.mod"}})`, cli.RepoName)
	requestBody := bytes.NewBufferString(reqB)

	req, err := http.NewRequest("POST", url, requestBody)
	if err != nil {
		log.Errorf("Create request err, %s", err)
		return nil
	}

	// Set username and password
	req.Header.Set("Content-Type", "text/plain")
	req.SetBasicAuth(cli.Username, cli.Password)

	// Do request
	resp, err := prhttp.DoHttpRequest(req)
	if err != nil {
		log.Errorf("fail to send req, %s", err)
		return nil
	}
	defer resp.Body.Close()

	bodyBytes, _ := io.ReadAll(resp.Body)

	type FileInfo struct {
		Repo       string `json:"repo"`
		Path       string `json:"path"`
		Name       string `json:"name"`
		Type       string `json:"type"`
		Size       int    `json:"size"`
		Created    string `json:"created"`
		CreatedBy  string `json:"created_by"`
		Modified   string `json:"modified"`
		ModifiedBy string `json:"modified_by"`
		Updated    string `json:"updated"`
	}

	var result struct {
		Results []FileInfo `json:"results"`
	}

	metaArr := []handler.ComponentView{}

	// 解析JSON字符串
	json.Unmarshal(bodyBytes, &result)

	// 遍历结果并打印文件名
	for _, file := range result.Results {
		fmt.Printf("%s\n", file.Path)

		mata := handler.ComponentView{
			Path: file.Path,
		}
		metaArr = append(metaArr, mata)
	}
	return metaArr
}
