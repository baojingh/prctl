package debian

import "github.com/spf13/cobra"

/**
prctl config list
prctl config add    --repo xxx-aa-sss
prctl config remove --repo xxx-aa-sss
**/

var ConfigCmd = &cobra.Command{
	Use:   "config",
	Short: "config repo for the session",
	Run: func(cmd *cobra.Command, args []string) {
		log.Info("config")
	},
}
