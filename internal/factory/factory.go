package factory

import (
	"github.com/baojingh/prctl/internal/deb"
	"github.com/baojingh/prctl/internal/golang"
	"github.com/baojingh/prctl/internal/handler"
	"github.com/baojingh/prctl/internal/pypi"
)

func NewRepoManageFactory(clientType string) handler.RepoManage {
	switch clientType {
	case "pypi":
		return pypi.NewPypiRepository()
	case "deb":
		return deb.NewDebRepository()
	case "go":
		return golang.NewGoRepository()
	default:
		return nil
	}

}
