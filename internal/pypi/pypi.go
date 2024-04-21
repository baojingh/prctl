package pypi

import (
	"github.com/baojingh/prctl/internal/common"
	"github.com/baojingh/prctl/internal/handler"
	"github.com/baojingh/prctl/internal/logger"
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
	log.Infof("pypi delete all, %v, %v", param, cli)

	// if param.All {
	// 	log.Infof("delete all, %v", param)
	// } else if param.Name != "" && param.Version == "" {
	// 	log.Infof("delete component, %v", param)
	// } else if param.Name != "" && param.Version != "" {
	// 	log.Infof("delete component and version, %v", param)
	// }
	// return nil
}

// input: /xx/xx/xx/ss.txt, check is it exists
// output aa/ss/ created if not exist
func (cli *PypiRepoManage) Download(input string, output string) {
	log.Info("pypi download")
}

func (cli *PypiRepoManage) Upload(meta handler.ComponentMeta, input string) {
	log.Infof("pypi upload")
}
