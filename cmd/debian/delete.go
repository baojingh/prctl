package debian

import (
	"github.com/baojingh/prctl/internal/deb"
	"github.com/spf13/cobra"
)

// prctl delete --all
// prctl delete --component aa --version 1.1
// prctl delete --component aa

var DeleteCmd = &cobra.Command{
	Use:   "delete",
	Short: "delete components in repo",
	Run: func(cmd *cobra.Command, args []string) {
		cmd.Flags().GetString("input")
		deb.DeleteDeb()
	},
}

func init() {
	DownloadCmd.Flags().StringP("all", "a", "", "Delete all components")
	DownloadCmd.Flags().StringP("component", "c", "", "Components name")
	DownloadCmd.Flags().StringP("version", "v", "", "Version")
}
