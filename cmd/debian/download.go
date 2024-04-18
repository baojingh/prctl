package debian

import (
	"github.com/baojingh/prctl/internal/deb"
	"github.com/spf13/cobra"
)

// prctl download --input components.txt --output ./outputs

var DownloadCmd = &cobra.Command{
	Use:   "download",
	Short: "download .deb files into local path",
	Run: func(cmd *cobra.Command, args []string) {
		input, _ := cmd.Flags().GetString("input")
		output, _ := cmd.Flags().GetString("output")
		cli := deb.NewClientFromConfig()
		if cli != nil {
			cli.Download(input, output)
		}
	},
}

func init() {
	DownloadCmd.Flags().StringP("input", "i", "", "Components list file")
	DownloadCmd.Flags().StringP("output", "o", "", "Components downloaded in to the directory")
	DownloadCmd.MarkFlagRequired("input")
	DownloadCmd.MarkFlagRequired("output")
}
