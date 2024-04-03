package debian

import (
	"github.com/baojingh/prctl/logger"
	"github.com/spf13/cobra"
)

var log = logger.New()

var LoginCmd = &cobra.Command{
	Use:   "login",
	Short: "login the repo",
	Run: func(cmd *cobra.Command, args []string) {
		url, _ := cmd.Flags().GetString("url")
		username, _ := cmd.Flags().GetString("username")
		password, _ := cmd.Flags().GetString("password")
		log.Infof("This is a login test, #%s#%s#%s", url, username, password)
		// login(url, username, password)
	},
}

func init() {
	LoginCmd.MarkFlagRequired("url")
	LoginCmd.MarkFlagRequired("username")
	LoginCmd.MarkFlagRequired("password")
}

// func login(url string, username string, password string) {
// 	log.Infof("This is a login test, #%s#%s#%s", url, username, password)
// }
