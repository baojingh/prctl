package repo

import (
	"os"

	"github.com/baojingh/prctl/cmd/factory"
	"github.com/spf13/cobra"
)

var LoginCmd = factory.CreateLoginCmd()
var LogoutCmd = factory.CreateLogioutmd()

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

func DebCmd() {
	DeleteCmd := factory.CreateDeleteCmd()
	DownloadCmd := factory.CreateDownloadCmd()
	UploadCmd := factory.CreateUploadCmd()
	DebianCommand.AddCommand(DownloadCmd)
	DebianCommand.AddCommand(UploadCmd)
	DebianCommand.AddCommand(DeleteCmd)
}

func PypiCmd() {
	DeleteCmd := factory.CreateDeleteCmd()
	DownloadCmd := factory.CreateDownloadCmd()
	UploadCmd := factory.CreateUploadCmd()
	PypiCommand.AddCommand(DownloadCmd)
	PypiCommand.AddCommand(UploadCmd)
	PypiCommand.AddCommand(DeleteCmd)
}

func init() {
	PypiCmd()
	DebCmd()
}
