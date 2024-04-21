package common

import (
	"github.com/baojingh/prctl/internal/factory"
	"github.com/spf13/cobra"
)

var DeleteCmd = &cobra.Command{
	Use:   "delete",
	Short: "delete components in repo",
	Run: func(cmd *cobra.Command, args []string) {
		// all, _ := cmd.Flags().GetBool("all")

		parentType := cmd.Parent().Use
		cli := factory.NewRepoManageFactory(parentType)
		if cli != nil {
			cli.Delete("all")
		}
	},
}

var DownloadCmd = &cobra.Command{
	Use:   "download",
	Short: "download files into local path",
	Run: func(cmd *cobra.Command, args []string) {
		// input, _ := cmd.Flags().GetString("input")
		// output, _ := cmd.Flags().GetString("output")
		parentType := cmd.Parent().Use
		cli := factory.NewRepoManageFactory(parentType)
		if cli != nil {
			cli.Download("download")
		}
	},
}

var UploadCmd = &cobra.Command{
	Use:   "upload",
	Short: "upload components into repo",
	Run: func(cmd *cobra.Command, args []string) {
		// distribution, _ := cmd.Flags().GetString("distribution")
		// component, _ := cmd.Flags().GetString("component")
		// input, _ := cmd.Flags().GetString("input")
		// architecture, _ := cmd.Flags().GetString("architecture")

		parentType := cmd.Parent().Use
		cli := factory.NewRepoManageFactory(parentType)
		if cli != nil {
			cli.Upload("upload")
		}
	},
}

func init() {

	DeleteCmd.Flags().Bool("all", false, "Delete all components")
	DeleteCmd.MarkFlagRequired("all")

	UploadCmd.Flags().StringP("distribution", "d", "", "distribution")
	UploadCmd.Flags().StringP("component", "c", "", "component")
	UploadCmd.Flags().StringP("architecture", "a", "", "architecture")
	UploadCmd.Flags().StringP("input", "i", "", "path to .deb")
	UploadCmd.MarkFlagRequired("distribution")
	UploadCmd.MarkFlagRequired("component")
	UploadCmd.MarkFlagRequired("architecture")
	UploadCmd.MarkFlagRequired("input")

	DownloadCmd.Flags().StringP("input", "i", "", "Components list file")
	DownloadCmd.Flags().StringP("output", "o", "", "Components downloaded in to the directory")
	DownloadCmd.MarkFlagRequired("input")
	DownloadCmd.MarkFlagRequired("output")
}
