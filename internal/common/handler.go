package handler

type ClientFactory interface {
	CreateClient() interface{}
}

type ClientOperations interface {
	Delete(param DeleteParam)
	Download(input string, output string)
	pload(meta ComponentMeta, input string)
}

func NewClientFactory(t string) ClientFactory {
	switch t {
	case "deb":
		return &DebClientFactory{}
	case "pypi":
		return &PypiClientFactory{}
	case "go":
		return &GoClientFactory{}
	default:
		return nil
	}
}
