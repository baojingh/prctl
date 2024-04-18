package deb

type DeleteParam struct {
	All     bool `json:"all" default:"false"`
	Name    string
	Version string
}

func (cli *Client) Delete(param DeleteParam) {
	if param.All {
		log.Infof("delete all, %v", param)
	} else if param.Name != "" && param.Version == "" {
		log.Infof("delete component, %v", param)
	} else if param.Name != "" && param.Version != "" {
		log.Infof("delete component and version, %v", param)

	}
}
