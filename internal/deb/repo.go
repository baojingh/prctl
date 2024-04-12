package deb

import (
	"encoding/base64"
	"encoding/json"
	"os"
	"os/user"
	"path/filepath"

	"github.com/baojingh/prctl/internal/utils/files"
)

type CredentialInfo struct {
	Url      string `json:"url"`
	Username string `json:"username"`
	Password string `json:"password"`
}

func ReadCred(credPath string) CredentialInfo {
	content, err := os.ReadFile(credPath)
	if err != nil {
		log.Errorf("Failed to open file: %s", err)
	}
	var credInfo CredentialInfo
	decodeCred, err := base64.StdEncoding.DecodeString(string(content))
	if err != nil {
		log.Errorf("Fail to decode, %s", err)
		return credInfo
	}
	err = json.Unmarshal(decodeCred, &credInfo)
	if err != nil {
		log.Errorf("Fail to unmarshal from decode data, %s", err)
		return credInfo
	}
	return credInfo
}

func WriteCred(cred CredentialInfo) error {
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
	credByte, err := json.Marshal(cred)
	if err != nil {
		log.Errorf("Cannot marshal struct to byte array, %s", err)
		return err
	}
	encodeCred := base64.StdEncoding.EncodeToString(credByte)
	// WriteFile create it if not exist
	credPath := filepath.Join(hiddenPath, ".config")
	err = os.WriteFile(credPath, []byte(encodeCred), 0600)
	if err != nil {
		log.Errorf("Cannot create credential path %s, %s", credPath, err)
		return err
	}
	log.Infof("Create config info success, %s", credPath)
	return nil
}
