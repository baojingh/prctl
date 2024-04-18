package debian

import (
	"github.com/baojingh/prctl/internal/deb"
	"github.com/baojingh/prctl/internal/logger"
	"github.com/spf13/cobra"
)

var log = logger.New()

var LoginCmd = &cobra.Command{
	Use:   "login",
	Short: "login the repo",
	Long:  ``,
	Run: func(cmd *cobra.Command, args []string) {
		url, _ := cmd.Flags().GetString("url")
		repo, _ := cmd.Flags().GetString("repo")
		username, _ := cmd.Flags().GetString("username")
		password, _ := cmd.Flags().GetString("password")
		deb.WriteCred(url, repo, username, password)
	},
}

func init() {
	LoginCmd.Flags().String("url", "", "URL of the repository")
	LoginCmd.Flags().String("repo", "", "name of the repository")
	LoginCmd.Flags().StringP("username", "u", "", "Username for login")
	LoginCmd.Flags().StringP("password", "p", "", "Password for login")
	// MarkFlagRequired is put in first place, then it will not take effect.
	// MarkFlagRequired must be in the second place
	LoginCmd.MarkFlagRequired("url")
	LoginCmd.MarkFlagRequired("repo")
	LoginCmd.MarkFlagRequired("username")
	LoginCmd.MarkFlagRequired("password")
}
