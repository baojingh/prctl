package handler

type DeleteParam struct {
	All     bool `json:"all" default:"false"`
	Name    string
	Version string
}

type ComponentMeta struct {
	Distribution string
	Component    string
	Architech    string
	Name         string
	Version      string
}

type Client struct {
	RepoUrl    string `json:"repoUrl"`
	RepoName   string `json:"repoName"`
	Username   string `json:"username"`
	Password   string `json:"password"`
	ConfigPath string `json:"configPath"`
}

type RepoManage interface {
	Delete(param string)
	Download(input string, output string)
	Upload(param string)
}
