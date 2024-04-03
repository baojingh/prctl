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
