package repo

import (
	"os"

	"github.com/baojingh/prctl/cmd/common"
	"github.com/baojingh/prctl/internal/cred"
	"github.com/spf13/cobra"
)

var LoginCmd = &cobra.Command{
	Use:   "login",
	Short: "login the repo",
	Long:  ``,
	Run: func(cmd *cobra.Command, args []string) {
		url, _ := cmd.Flags().GetString("url")
		repo, _ := cmd.Flags().GetString("repo")
		username, _ := cmd.Flags().GetString("username")
		password, _ := cmd.Flags().GetString("password")
		cred.Login(url, repo, username, password)
	},
}

var LogoutCmd = &cobra.Command{
	Use:   "logout",
	Short: "logout the repo",
	Long:  ``,
	Run: func(cmd *cobra.Command, args []string) {
		cred.Logout()
	},
}

var DebianCommand = &cobra.Command{
	Use:   "deb",
	Short: "Process deb component in repo",
	Run: func(cmd *cobra.Command, args []string) {
		cmd.Help()
		os.Exit(1)
	},
}

var PypiCommand = &cobra.Command{
	Use:   "pypi",
	Short: "Process pypi component in python repo",
	Run: func(cmd *cobra.Command, args []string) {
		cmd.Help()
		os.Exit(1)
	},
}

var GoCommand = &cobra.Command{
	Use:   "go",
	Short: "Process go component in go repo",
	Run: func(cmd *cobra.Command, args []string) {
		cmd.Help()
		os.Exit(1)
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

	PypiCommand.AddCommand(common.DownloadCmd)
	PypiCommand.AddCommand(common.UploadCmd)
	PypiCommand.AddCommand(common.DeleteCmd)

	DebianCommand.AddCommand(common.DownloadCmd)
	DebianCommand.AddCommand(common.UploadCmd)
	DebianCommand.AddCommand(common.DeleteCmd)

	GoCommand.AddCommand(common.DownloadCmd)
	GoCommand.AddCommand(common.UploadCmd)
	GoCommand.AddCommand(common.DeleteCmd)

}
