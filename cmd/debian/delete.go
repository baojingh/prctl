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
		param := deb.DeleteParam{
			All: all,
		}
		cli := deb.NewClientFromConfig()
		if cli != nil {
			cli.Delete(param)
		}
	},
}

func init() {
	DeleteCmd.Flags().Bool("all", false, "Delete all components")
	DeleteCmd.MarkFlagRequired("all")
}
