package pypi

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"os"
	"path/filepath"
	"regexp"
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

type PypiRepoManage struct {
	handler.Client
}

var log = logger.New()

func NewPypiRepository() handler.RepoManage {
	cli := common.CreateClient()
	return &PypiRepoManage{Client: *cli}
}

// curl  -u cs:cds  -X DELETE https://cs.rtf-alm.cs.cloud/artifactory/cs-dev-pypi-awsl/PyYAML
func (cli *PypiRepoManage) Delete(param handler.DeleteParam) {
	metaArr := cli.List()

	for _, ele := range metaArr {
		// Create request object
		name := strings.Split(ele.Name, "-")[0]
		url := fmt.Sprintf("%s/artifactory/%s/%s", cli.RepoUrl, cli.RepoName, name)

		req, err := http.NewRequest("DELETE", url, nil)
		if err != nil {
			log.Errorf("Create request err, %s", err)
			return
		}

		// Set username and password
		req.Header.Set("Content-Type", "text/plain")
		req.SetBasicAuth(cli.Username, cli.Password)

		// Do request
		prhttp.DoHttpRequest(req)
	}
	url := fmt.Sprintf("%s/artifactory/%s/%s", cli.RepoUrl, cli.RepoName, ".pypi")

	req, err := http.NewRequest("DELETE", url, nil)
	if err != nil {
		log.Errorf("Create request err, %s", err)
		return
	}

	// Set username and password
	req.Header.Set("Content-Type", "text/plain")
	req.SetBasicAuth(cli.Username, cli.Password)

	// Do request
	prhttp.DoHttpRequest(req)

}

// input: /xx/xx/xx/ss.txt, check is it exists
// output aa/ss/ created  by pip command automatically
func (cli *PypiRepoManage) Download(input string, output string) {
	// pip download -d sample/vcas/vca  --only-binary=:all:  -r ./requirements.txt
	params := []string{"download", "-d", output, "--only-binary=:all: ", "-r", input}
	log.Infof("Command: pip %s", strings.Join(params, " "))
	out, err := shell.DoShellCmd("pip", params...)
	if err != nil {
		log.Errorf("failed to download Pypi Components, err: %v", out)
		return
	}
	log.Info("Download Pypi success.")
}

func (cli *PypiRepoManage) Upload(meta handler.ComponentMeta, input string) {
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
func (cli *PypiRepoManage) doUpload(path string, fileName string) {

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

func (cli *PypiRepoManage) List() []handler.ComponentView {
	// Create request object
	url := fmt.Sprintf("%s/%s", cli.RepoUrl, "artifactory/api/search/aql")

	reqB := fmt.Sprintf(`items.find({"repo":"%s"})`, cli.RepoName)
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
	re := regexp.MustCompile("^.pypi.*")

	// 遍历结果并打印文件名
	for _, file := range result.Results {
		if re.MatchString(file.Path) {
			continue
		}
		fmt.Printf("%s %s\n", file.Path, file.Name)

		mata := handler.ComponentView{
			Name:     strings.Split(file.Path, "/")[0],
			Version:  strings.Split(file.Path, "/")[1],
			FileName: file.Name,
			Path:     file.Path,
		}
		metaArr = append(metaArr, mata)
	}
	return metaArr
}
