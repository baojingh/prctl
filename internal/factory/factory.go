package factory

import (
	"github.com/baojingh/prctl/internal/deb"
	"github.com/baojingh/prctl/internal/handler"
	"github.com/baojingh/prctl/internal/pypi"
)

type ClientFactory interface {
	CreateClient() handler.ClientOperations
}

func NewClientFactory(clientType string) ClientFactory {
	switch clientType {
	case "pypi":
		return &pypi.PypiClientFactory{}
	case "deb":
		return &deb.DebClientFactory{}
	default:
		return nil
	}

}
