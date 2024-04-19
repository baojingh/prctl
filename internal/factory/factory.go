package factory

import (
	"github.com/baojingh/prctl/internal/deb"
	"github.com/baojingh/prctl/internal/handler"
	"github.com/baojingh/prctl/internal/pypi"
)

func NewRepoManageFactory(clientType string) handler.RepoManage {
	switch clientType {
	case "pypi":
		return pypi.NewPypiRepository()
	case "deb":
		return deb.NewDebRepository()
	default:
		return nil
	}

}
