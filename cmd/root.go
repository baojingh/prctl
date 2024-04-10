package cmd

import (
	"fmt"
	"os"

	"github.com/baojingh/prctl/cmd/debian"
	"github.com/spf13/cobra"
)

var (
	buildTime    string
	buildVersion string
)

var RootCmd = &cobra.Command{
	Use:     "prctl",
	Short:   "prctl update components in private repo, such as debian/golang/pypi type in Jfrog/Nexus",
	Long:    `prctl update components in private repo, such as debian/golang/pypi type in JFrog/Nexus`,
	Version: fmt.Sprintf("%s, build in %s", buildVersion, buildTime),
	Run: func(cmd *cobra.Command, args []string) {
		cmd.Help()
		os.Exit(0)
	},
}

func init() {
	RootCmd.AddCommand(debian.LoginCmd)
	RootCmd.AddCommand(debian.LogoutCmd)
	RootCmd.AddCommand(debian.DownloadCmd)
}
