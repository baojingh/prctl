package debian

import (
	"github.com/baojingh/prctl/pkg/files"
	"github.com/spf13/cobra"
)

var LogoutCmd = &cobra.Command{
	Use:   "logout",
	Short: "logout the repo",
	Long:  ``,
	Run: func(cmd *cobra.Command, args []string) {
		path := "/home/ubuntu/.prctl/.config"
		logout(path)
	},
}

func logout(path string) {
	success := files.RemoveFileIfExist(path)
	if success {
		log.Infof("Remove cred file success, %s", path)
	} else {
		log.Debugf("Cred file not exists in %s", path)
	}
	log.Info("Logout success.")
}
