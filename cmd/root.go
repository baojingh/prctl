package cmd

import (
	"fmt"
	"github.com/baojingh/prctl/cmd/cred"
	"github.com/baojingh/prctl/cmd/debian"
	"github.com/baojingh/prctl/cmd/pypi"
	"os"

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
	RootCmd.AddCommand(cred.LoginCmd)
	RootCmd.AddCommand(cred.LogoutCmd)

	RootCmd.AddCommand(debian.DebianCommand)
	RootCmd.AddCommand(pypi.PypiCommand)
}
