package debian

import (
	"os"

	"github.com/baojingh/prctl/internal/deb"
	"github.com/spf13/cobra"
)

var DebianCommand = &cobra.Command{
	Use:   "deb",
	Short: "Process deb component in repo",
	Run: func(cmd *cobra.Command, args []string) {
		cmd.Help()
		os.Exit(1)
	},
}

var DeleteCmd = &cobra.Command{
	Use:   "delete",
	Short: "delete components in repo",
	Run: func(cmd *cobra.Command, args []string) {
		all, _ := cmd.Flags().GetBool("all")
		param := deb.DeleteParam{
			All: all,
		}

		// parentType := cmd.Parent().Use

		cli := deb.NewClientFromConfig()
		if cli != nil {
			cli.Delete(param)
		}
	},
}

var DownloadCmd = &cobra.Command{
	Use:   "download",
	Short: "download files into local path",
	Run: func(cmd *cobra.Command, args []string) {
		input, _ := cmd.Flags().GetString("input")
		output, _ := cmd.Flags().GetString("output")
		cli := deb.NewClientFromConfig()
		if cli != nil {
			cli.Download(input, output)
		}
	},
}

var UploadCmd = &cobra.Command{
	Use:   "upload",
	Short: "upload components into repo",
	Run: func(cmd *cobra.Command, args []string) {
		distribution, _ := cmd.Flags().GetString("distribution")
		component, _ := cmd.Flags().GetString("component")
		input, _ := cmd.Flags().GetString("input")
		architecture, _ := cmd.Flags().GetString("architecture")
		meta := deb.ComponentMeta{
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
	DebianCommand.AddCommand(DownloadCmd)
	DebianCommand.AddCommand(UploadCmd)
	DebianCommand.AddCommand(DeleteCmd)

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
