package debian

import (
	"github.com/baojingh/prctl/internal/deb"
	"github.com/spf13/cobra"
)

/**
prctl upload   --type debian  --path /var/canbu/vwnui
**/

var UploadCmd = &cobra.Command{
	Use:   "upload",
	Short: "upload debian components into repo",
	Run: func(cmd *cobra.Command, args []string) {
		distribution, _ := cmd.Flags().GetString("distribution")
		component, _ := cmd.Flags().GetString("component")
		input, _ := cmd.Flags().GetString("input")
		architecture, _ := cmd.Flags().GetString("architecture")
		meta := deb.DebComponentMeta{
			Distribution: distribution,
			Component:    component,
			Architech:    architecture,
		}
		cli := deb.NewClientFromConfig()
		if cli != nil {
			cli.Upload(meta, input)
		}
	},
}

func init() {
	UploadCmd.Flags().StringP("distribution", "d", "", "distribution")
	UploadCmd.Flags().StringP("component", "c", "", "component")
	UploadCmd.Flags().StringP("architecture", "a", "", "architecture")
	UploadCmd.Flags().StringP("input", "i", "", "path to .deb")
	UploadCmd.MarkFlagRequired("distribution")
	UploadCmd.MarkFlagRequired("component")
	UploadCmd.MarkFlagRequired("architecture")
	UploadCmd.MarkFlagRequired("input")
}
