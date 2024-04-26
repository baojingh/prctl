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
	ListCmd := factory.CreateListCmd()

	DebianCommand.AddCommand(DownloadCmd)
	DebianCommand.AddCommand(UploadCmd)
	DebianCommand.AddCommand(DeleteCmd)
	DebianCommand.AddCommand(ListCmd)
}

func PypiCmd() {
	DeleteCmd := factory.CreateDeleteCmd()
	DownloadCmd := factory.CreateDownloadCmd()
	UploadCmd := factory.CreateUploadCmd()
	ListCmd := factory.CreateListCmd()
	PypiCommand.AddCommand(DownloadCmd)
	PypiCommand.AddCommand(UploadCmd)
	PypiCommand.AddCommand(DeleteCmd)
	PypiCommand.AddCommand(ListCmd)
}

func init() {
	PypiCmd()
	DebCmd()
}
