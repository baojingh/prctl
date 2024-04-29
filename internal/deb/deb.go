package deb

import (
	"bufio"
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

type DebRepoManage struct {
	handler.Client
}

func NewDebRepository() handler.RepoManage {
	cli := common.CreateClient()
	return &DebRepoManage{Client: *cli}
}

var log = logger.New()

func (j *DebRepoManage) Delete(param handler.DeleteParam) {
	log.Infof("debian delete all, %v, %v", param, j)
}

// input: /xx/xx/xx/ss.txt, check is it exists
// output aa/ss/ created if not exist
func (cli *DebRepoManage) Download(input string, output string) {
	// init debian environment
	_, err := shell.DoShellCmd("apt-get", "update")
	if err != nil {
		log.Errorf("fail to apt-get update deb env, err: %s", err)
		return
	}
	file, err := os.Open(input)
	if err != nil {
		log.Errorf("Cannot open file %s", input)
		return
	}
	defer file.Close()

	//  Create the output dir if it not exist
	files.CreateDirIfNotExist(output, 0755)

	var buffer strings.Builder
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		name := scanner.Text()
		buffer.WriteString(name)
		buffer.WriteString(" ")
	}
	if err := scanner.Err(); err != nil {
		return
	}
	res := strings.TrimSpace(buffer.String())
	changeDirAndDo(res, output)
	log.Info("Deb components are downloaded success.")
}

// https://stackoverflow.com/questions/52435908/how-to-change-the-shells-current-working-directory-in-go
// apt-get download just put the components in current path, so we need change to target dir
func changeDirAndDo(nameList string, path string) {
	cwd, _ := os.Getwd()
	if err := os.Chdir(path); err != nil {
		return
	}
	// component name list must be seperated and then composed by append.
	params := []string{"download"}
	params = append(params, strings.Fields(nameList)...)
	log.Infof("Command: apt-get %s", strings.Join(params, " "))
	out, err := shell.DoShellCmd("apt-get", params...)
	if err != nil {
		log.Errorf("failed to download %s, err: %s, out: %s", nameList, err, out)
		return
	}
	log.Infof("Download %s success.", nameList)

	if err := os.Chdir(cwd); err != nil {
		return
	}
}

func (cli *DebRepoManage) Upload(meta handler.ComponentMeta, input string) {
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
func (cli *DebRepoManage) doUpload(meta handler.ComponentMeta, path string, fileName string) {
	arch := meta.Architech
	dis := meta.Distribution
	com := meta.Component

	uploadUrl := fmt.Sprintf("%s/artifactory/%s/pool/%s;deb.distribution=%s;deb.component=%s;deb.architecture=%s",
		cli.RepoUrl, cli.RepoName, fileName, dis, com, arch)

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

// curl  -u fds:fs  -XGET https://fa.rtf-alm.
// fsa.cloud/artifactory/api/storage/fa-dev-debian-awsl/pool?list
func (cli *DebRepoManage) List() []handler.ComponentView {
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
