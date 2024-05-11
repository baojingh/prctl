package factory

import (
	"github.com/baojingh/prctl/internal/cred"
	"github.com/baojingh/prctl/internal/factory"
	"github.com/baojingh/prctl/internal/handler"
	"github.com/spf13/cobra"
)

func CreateDeleteCmd() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "delete",
		Short: "delete components in repo",
		Run: func(cmd *cobra.Command, args []string) {
			all, _ := cmd.Flags().GetBool("all")

			parentType := cmd.Parent().Use
			cli := factory.NewRepoManageFactory(parentType)
			if cli != nil {
				cli.Delete(handler.DeleteParam{All: all})
			}
		},
	}
	cmd.Flags().Bool("all", false, "Delete all components")
	cmd.MarkFlagRequired("all")
	return cmd
}

func CreateDownloadCmd() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "download",
		Short: "download files into local path",
		Run: func(cmd *cobra.Command, args []string) {
			input, _ := cmd.Flags().GetString("input")
			output, _ := cmd.Flags().GetString("output")
			parentType := cmd.Parent().Use
			cli := factory.NewRepoManageFactory(parentType)
			if cli != nil {
				cli.Download(input, output)
			}
		},
	}

	cmd.Flags().StringP("input", "i", "", "Components list file")
	cmd.Flags().StringP("output", "o", "", "Components downloaded in to the directory")
	cmd.MarkFlagRequired("input")
	cmd.MarkFlagRequired("output")

	return cmd

}

func CreateUploadCmd() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "upload",
		Short: "upload components into repo",
		Run: func(cmd *cobra.Command, args []string) {
			distribution, _ := cmd.Flags().GetString("distribution")
			component, _ := cmd.Flags().GetString("component")
			input, _ := cmd.Flags().GetString("input")
			architecture, _ := cmd.Flags().GetString("architecture")

			parentType := cmd.Parent().Use
			cli := factory.NewRepoManageFactory(parentType)
			if cli != nil {
				cli.Upload(
					handler.ComponentMeta{
						Distribution: distribution,
						Component:    component,
						Architech:    architecture,
					},
					input)
			}
		},
	}
	cmd.Flags().StringP("input", "i", "", "path to .deb")
	cmd.MarkFlagRequired("input")

	return cmd
}

func CreateLoginCmd() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "login",
		Short: "login the repo",
		Long:  ``,
		Run: func(cmd *cobra.Command, args []string) {
			url, _ := cmd.Flags().GetString("url")
			repo, _ := cmd.Flags().GetString("repo")
			username, _ := cmd.Flags().GetString("username")
			password, _ := cmd.Flags().GetString("password")
			cred.Login(url, repo, username, password)
		},
	}

	cmd.Flags().String("url", "", "URL of the repository")
	cmd.Flags().String("repo", "", "name of the repository")
	cmd.Flags().StringP("username", "u", "", "Username for login")
	cmd.Flags().StringP("password", "p", "", "Password for login")
	// MarkFlagRequired is put in first place, then it will not take effect.
	// MarkFlagRequired must be in the second place
	cmd.MarkFlagRequired("url")
	cmd.MarkFlagRequired("repo")
	cmd.MarkFlagRequired("username")
	cmd.MarkFlagRequired("password")
	return cmd
}

func CreateLogioutmd() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "logout",
		Short: "logout the repo",
		Long:  ``,
		Run: func(cmd *cobra.Command, args []string) {
			cred.Logout()
		},
	}
	return cmd
}

func CreateListCmd() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "list",
		Short: "logout the repo",
		Long:  ``,
		Run: func(cmd *cobra.Command, args []string) {
			cmd.Flags().GetBool("all")
			parentType := cmd.Parent().Use
			cli := factory.NewRepoManageFactory(parentType)
			if cli != nil {
				cli.List()
			}
		},
	}

	cmd.Flags().Bool("all", false, "List all components")
	cmd.MarkFlagRequired("all")

	return cmd
}
