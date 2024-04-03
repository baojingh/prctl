package debian

import (
	"encoding/base64"
	"fmt"
	"os"
	"os/user"
	"path/filepath"

	"github.com/baojingh/prctl/logger"
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

func login(url string, username string, password string) {
	log.Infof("This is a login test, #%s#%s#%s", url, username, password)
	currUser, err := user.Current()
	if err != nil {
		log.Fatalf("cannot get current user default path, %s", err)
		return
	}
	userPath := currUser.HomeDir
	hiddenPath := filepath.Join(userPath, ".prctl")
	info, err := os.Stat(hiddenPath)
	if !os.IsNotExist(err) {
		err = os.Mkdir(hiddenPath, 0700)
		if err != nil {
			log.Fatalf("Cannot create hidden path, %s", err)
			return
		}
	} else {
		log.Info(info.IsDir())
	}

	cred := fmt.Sprintf("%s:%s", username, password)
	encodeCred := base64.StdEncoding.EncodeToString([]byte(cred))
	credPath := filepath.Join(hiddenPath, ".config")
	err = os.WriteFile(credPath, []byte(encodeCred), 0600)

	if err != nil {
		log.Fatalf("Cannot create credential file, %s", err)
		return
	}

	log.Infof("Create config info success, %s", credPath)
}
