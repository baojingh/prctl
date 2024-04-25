package pypi

import (
	"io"
	"net/http"

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
	url := "https://falcon.rtf-alm.siemens.cloud:443/artifactory/xo_cys-dev-pypi-awsl?list"
	// Create request object
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		log.Error(err)
		return nil
	}

	// Set username and password
	req.Header.Set("Content-Type", "application/octet-stream")
	req.SetBasicAuth("dd", "dd")

	// Do request
	resp, err := prhttp.DoHttpRequest(req)
	if err != nil {
		log.Error(err)
		return nil
	}
	defer resp.Body.Close()

	bodyBytes, _ := io.ReadAll(resp.Body)
	log.Infof("%v", bodyBytes)

	// // 将字节切片解码为struct对象
	// var myStruct MyStruct
	// _ = json.Unmarshal(bodyBytes, &myStruct)

	// // 打印struct对象的内容
	// log.Infof("Field1: %s, Field2: %d\n", myStruct.Field1, myStruct.Field2)

	return nil
}
