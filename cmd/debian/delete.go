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
		all, _ := cmd.Flags().GetBool("all")
		name, _ := cmd.Flags().GetString("name")
		version, _ := cmd.Flags().GetString("version")
		param := deb.DeleteParam{
			All:     all,
			Name:    name,
			Version: version,
		}
		deb.DeleteDeb(param)
	},
}

func init() {
	// TODO: delete components; delete component with version.
	// Now: Implement delete all components at present.[20240413]
	DownloadCmd.Flags().Bool("all", false, "Delete all components")
	DownloadCmd.Flags().StringP("name", "n", "", "Components name")
	DownloadCmd.Flags().StringP("version", "v", "", "Version")
}
