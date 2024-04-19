package handler

type ClientFactory interface {
	CreateClient() interface{}
}

func NewClientFactory(clientType string) ClientFactory {
	switch clientType {
	case "pypi":
		return &PypiClientFactory{}
	case "deb":
		return &DebClientFactory{}
	default:
		return nil
	}

}
