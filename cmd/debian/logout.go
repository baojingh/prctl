package debian

import (
	"github.com/spf13/cobra"
)

var LogoutCmd = &cobra.Command{
	Use:   "logout",
	Short: "logout the repo",
	Long:  ``,
	Run: func(cmd *cobra.Command, args []string) {
		url, _ := cmd.Flags().GetString("url")
		username, _ := cmd.Flags().GetString("username")
		password, _ := cmd.Flags().GetString("password")
		logout(url, username, password)
	},
}

func logout(url string, username string, password string) {
	log.Infof("This is a login test, #%s#%s#%s", url, username, password)
}
