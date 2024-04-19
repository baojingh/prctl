package pypi

import (
	"encoding/base64"
	"encoding/json"
	"os"
	"path/filepath"

	"github.com/baojingh/prctl/internal/handler"
	"github.com/baojingh/prctl/internal/logger"
	"github.com/baojingh/prctl/pkg/prsys"
)

type DeleteParam struct {
	All     bool `json:"all" default:"false"`
	Name    string
	Version string
}

type ComponentMeta struct {
	Distribution string
	Component    string
	Architech    string
	Name         string
	Version      string
}

type Client struct {
	RepoUrl    string `json:"repoUrl"`
	RepoName   string `json:"repoName"`
	Username   string `json:"username"`
	Password   string `json:"password"`
	ConfigPath string `json:"configPath"`
}

type PypiRepoManage struct {
	Client
}

var log = logger.New()

func NewPypiRepository() handler.RepoManage {
	cli := CreateClient()
	return &PypiRepoManage{Client: *cli}
}

// get cred path, default os /home/${USER}/.prctl/config or /root/.prctl/config
func getConfigPath() string {
	userPath := prsys.CurrentUserPath()
	hiddenPath := filepath.Join(userPath, ".prctl")
	configPath := filepath.Join(hiddenPath, "config")
	return configPath
}

// Create client from a cred file, apply in: logout,download, upload, delete
func CreateClient() *Client {
	configPath := getConfigPath()
	content, err := os.ReadFile(configPath)
	if err != nil {
		log.Errorf("configPath %s not exist.", configPath)
		return nil
	}
	var cli Client
	decodeCred, err := base64.StdEncoding.DecodeString(string(content))
	if err != nil {
		log.Errorf("fail to decode, %s", err)
		return nil
	}
	err = json.Unmarshal(decodeCred, &cli)
	if err != nil {
		log.Errorf("fail to unmarshal from decode data, %s", err)
		return nil
	}
	return &cli
}

func (cli *PypiRepoManage) Delete(param string) {
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
func (cli *PypiRepoManage) Download(param string) {
	log.Info("pypi download")
}

func (cli *PypiRepoManage) Upload(param string) {
	log.Infof("pypi upload")
}
