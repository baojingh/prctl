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
	UploadCmd.Flags().String("distribution", "", "distribution")
	UploadCmd.Flags().String("component", "", "component")
	UploadCmd.Flags().String("architecture", "", "architecture")
	UploadCmd.Flags().StringP("input", "i", "", "architecture")
	UploadCmd.MarkFlagRequired("distribution")
	UploadCmd.MarkFlagRequired("component")
	UploadCmd.MarkFlagRequired("architecture")
	UploadCmd.MarkFlagRequired("input")
}
