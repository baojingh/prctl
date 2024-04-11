package debian

import (
	"github.com/spf13/cobra"
)

/**
prctl upload   --type debian  --path /var/canbu/vwnui
**/

var UploadCmd = &cobra.Command{
	Use:   "upload",
	Short: "upload debian components into repo",
	Run: func(cmd *cobra.Command, args []string) {
		log.Info("upload")
	},
}

func init() {
	UploadCmd.Flags().StringP("distribution", "dis", "", "distribution")
	UploadCmd.Flags().StringP("component", "com", "", "component")
	UploadCmd.Flags().StringP("architecture", "arch", "", "architecture")
	UploadCmd.Flags().StringP("input", "i", "", "architecture")
	UploadCmd.MarkFlagRequired("distribution")
	UploadCmd.MarkFlagRequired("component")
	UploadCmd.MarkFlagRequired("architecture")
	UploadCmd.MarkFlagRequired("input")
}
