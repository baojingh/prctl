package pypi

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"strings"

	"github.com/baojingh/prctl/internal/common"
	"github.com/baojingh/prctl/internal/handler"
	"github.com/baojingh/prctl/internal/logger"
	"github.com/baojingh/prctl/pkg/prhttp"
)

type PypiRepoManage struct {
	handler.Client
}

var log = logger.New()

func NewPypiRepository() handler.RepoManage {
	cli := common.CreateClient()
	return &PypiRepoManage{Client: *cli}
}

func (cli *PypiRepoManage) Delete(param handler.DeleteParam) {
	metaArr := cli.List()

	for _, ele := range metaArr {
		// Create request object
		url := fmt.Sprintf("%s/artifactory/%s/%s", cli.RepoUrl, cli.RepoName, ele.Name)

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

}

// input: /xx/xx/xx/ss.txt, check is it exists
// output aa/ss/ created if not exist
func (cli *PypiRepoManage) Download(input string, output string) {
	log.Info("pypi download")
}

func (cli *PypiRepoManage) Upload(meta handler.ComponentMeta, input string) {
	log.Infof("pypi upload")
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
	// 遍历结果并打印文件名
	for _, file := range result.Results {
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
