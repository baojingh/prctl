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

type PypiClient struct {
	handler.ClientOperations
}

type PypiClientFactory struct{}

var log = logger.New()

// Create client from a cred file, apply in: logout,download, upload, delete
func (j *PypiClientFactory) CreateClient() handler.ClientOperations {
	configPath := getConfigPath()
	content, err := os.ReadFile(configPath)
	if err != nil {
		log.Errorf("configPath %s not exist.", configPath)
		return nil
	}
	var cli PypiClient
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

// get cred path, default os /home/${USER}/.prctl/config or /root/.prctl/config
func getConfigPath() string {
	userPath := prsys.CurrentUserPath()
	hiddenPath := filepath.Join(userPath, ".prctl")
	configPath := filepath.Join(hiddenPath, "config")
	return configPath
}

func (cli *PypiClient) Delete(param string) {
	log.Infof("delete all, %v", param)

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
func (cli *PypiClient) Download(param string) {
	log.Info("pypi download")
}

func (cli *PypiClient) Upload(param string) {
	log.Infof("pypi upload")
}
