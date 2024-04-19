package handler

import (
	"encoding/base64"
	"encoding/json"
	"os"
	"path/filepath"

	"github.com/baojingh/prctl/internal/logger"
	"github.com/baojingh/prctl/pkg/files"
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

type PypiClientFactory struct{}

var log = logger.New()

// Create client from a cred file, apply in: logout,download, upload, delete
func (j *PypiClientFactory) CreateClient() interface{} {
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

// get cred path, default os /home/${USER}/.prctl/config or /root/.prctl/config
func getConfigPath() string {
	userPath := prsys.CurrentUserPath()
	hiddenPath := filepath.Join(userPath, ".prctl")
	configPath := filepath.Join(hiddenPath, "config")
	return configPath
}

func (cli *Client) Logout() {
	path := cli.ConfigPath
	success := files.RemoveFileIfExist(path)
	if success {
		log.Infof("Remove cred file success, %s", path)
	} else {
		log.Debugf("Config file not exists in %s", path)
	}
	log.Info("Logout success.")
}

func (cli *Client) Delete(param DeleteParam) error {
	if param.All {
		log.Infof("delete all, %v", param)
	} else if param.Name != "" && param.Version == "" {
		log.Infof("delete component, %v", param)
	} else if param.Name != "" && param.Version != "" {
		log.Infof("delete component and version, %v", param)
	}
	return nil
}

// input: /xx/xx/xx/ss.txt, check is it exists
// output aa/ss/ created if not exist
func (cli *Client) Download(input string, output string) {
	log.Info("Deb components are downloaded success.")
}

func (cli *Client) Upload(meta ComponentMeta, input string) {
	log.Infof("Response Status: %s\n", "")
}
