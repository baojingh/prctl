package debian

import (
	"encoding/base64"
	"encoding/json"
	"os"
	"os/user"
	"path/filepath"

	"github.com/baojingh/prctl/logger"
	"github.com/baojingh/prctl/utils"
	"github.com/spf13/cobra"
)

var log = logger.New()

var LoginCmd = &cobra.Command{
	Use:   "login",
	Short: "login the repo",
	Long:  ``,
	Run: func(cmd *cobra.Command, args []string) {
		url, _ := cmd.Flags().GetString("url")
		username, _ := cmd.Flags().GetString("username")
		password, _ := cmd.Flags().GetString("password")
		login(url, username, password)
	},
}

func init() {
	LoginCmd.MarkFlagRequired("url")
	LoginCmd.MarkFlagRequired("username")
	LoginCmd.MarkFlagRequired("password")
	LoginCmd.Flags().String("url", "", "URL of the repository")
	LoginCmd.Flags().String("username", "", "Username for login")
	LoginCmd.Flags().String("password", "", "Password for login")
}

type CredentialInfo struct {
	Url      string `json: url`
	Username string `json: username`
	Password string `json: password`
}

func login(url string, username string, password string) {
	cred := CredentialInfo{
		Url:      url,
		Username: username,
		Password: password,
	}
	WriteCred(cred)
}

func ReadCred(credPath string) CredentialInfo {
	var credInfo CredentialInfo
	decodeCred, err := base64.StdEncoding.DecodeString(credPath)
	if err != nil {
		log.Fatalf("Fail to decode, %s", err)
		return credInfo
	}
	err = json.Unmarshal(decodeCred, &credInfo)
	if err != nil {
		log.Fatalf("Fail to unmarshal from decode data, %s", err)
		return credInfo
	}
	return credInfo
}

func WriteCred(cred CredentialInfo) error {
	currUser, err := user.Current()
	if err != nil {
		log.Fatalf("cannot get current user default path, %s", err)
		return err
	}
	userPath := currUser.HomeDir
	hiddenPath := filepath.Join(userPath, ".prctl")
	err = utils.CreateDirIfNotExist(hiddenPath, 0700)
	if err != nil {
		log.Fatalf("Cannot create hidden path %s, %s", hiddenPath, err)
		return err
	}
	credByte, err := json.Marshal(cred)
	if err != nil {
		log.Fatalf("Cannot marshal struct to byte array, %s", err)
		return err
	}
	encodeCred := base64.StdEncoding.EncodeToString(credByte)
	// WriteFile create it if not exist
	credPath := filepath.Join(hiddenPath, ".config")
	err = os.WriteFile(credPath, []byte(encodeCred), 0600)
	if err != nil {
		log.Fatalf("Cannot create credential path %s, %s", credPath, err)
		return err
	}
	log.Infof("Create config info success, %s", credPath)
	return nil
}
