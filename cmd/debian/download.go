package debian

import "github.com/spf13/cobra"

// prctl download --input components.txt --output ./outputs

var DownloadCmd = &cobra.Command{
	Use:   "download",
	Short: "download .deb files into local path",
	Run: func(cmd *cobra.Command, args []string) {

	},
}
