package common

import (
	"encoding/base64"
	"encoding/json"
	"os"
	"path/filepath"

	"github.com/baojingh/prctl/internal/handler"
	"github.com/baojingh/prctl/internal/logger"
	"github.com/baojingh/prctl/pkg/prsys"
)

var log = logger.New()

// get cred path, default os /home/${USER}/.prctl/config or /root/.prctl/config
func GetConfigPath() string {
	userPath := prsys.CurrentUserPath()
	hiddenPath := filepath.Join(userPath, ".prctl")
	configPath := filepath.Join(hiddenPath, "config")
	return configPath
}

// Create client from a cred file, apply in: logout,download, upload, delete
func CreateClient() *handler.Client {
	configPath := GetConfigPath()
	content, err := os.ReadFile(configPath)
	if err != nil {
		log.Errorf("configPath %s not exist.", configPath)
		return nil
	}
	var cli handler.Client
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
