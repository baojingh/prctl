package deb

import (
	"encoding/base64"
	"encoding/json"
	"os"
	"os/user"
	"path/filepath"

	"github.com/baojingh/prctl/internal/logger"
	"github.com/baojingh/prctl/pkg/files"
	prsys "github.com/baojingh/prctl/pkg/prsys"
)

type Client struct {
	RepoUrl    string `json:"repoUrl"`
	RepoName   string `json:"repoName"`
	Username   string `json:"username"`
	Password   string `json:"password"`
	ConfigPath string `json:"configPath"`
}

var log = logger.New()

// Write cred info when logining
func WriteCred(url string, repo string, username string, password string) error {
	currUser, err := user.Current()
	if err != nil {
		log.Errorf("cannot get current user default path, %s", err)
		return err
	}
	userPath := currUser.HomeDir
	hiddenPath := filepath.Join(userPath, ".prctl")
	err = files.CreateDirIfNotExist(hiddenPath, 0700)
	if err != nil {
		log.Errorf("Cannot create hidden path %s, %s", hiddenPath, err)
		return err
	}
	// WriteFile create it if not exist
	configPath := filepath.Join(hiddenPath, "config")
	cli := Client{
		RepoUrl:    url,
		RepoName:   repo,
		Username:   username,
		Password:   password,
		ConfigPath: configPath,
	}
	credByte, err := json.Marshal(cli)
	if err != nil {
		log.Errorf("Cannot marshal struct to byte array, %s", err)
		return err
	}
	encodeCred := base64.StdEncoding.EncodeToString(credByte)

	err = os.WriteFile(configPath, []byte(encodeCred), 0600)
	if err != nil {
		log.Errorf("Cannot create credential path %s, %s", configPath, err)
		return err
	}
	log.Infof("Create config info success, %s", configPath)
	return nil
}

// get cred path, default os /home/${USER}/.prctl/config or /root/.prctl/config
func getConfigPath() string {
	userPath := prsys.CurrentUserPath()
	hiddenPath := filepath.Join(userPath, ".prctl")
	configPath := filepath.Join(hiddenPath, "config")
	return configPath
}

// Create client from a cred file, apply in: logout,download, upload, delete
func NewClientFromConfig() *Client {
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
