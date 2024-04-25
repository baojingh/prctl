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

type ComponentView struct {
	Name    string
	Version string
	Time    string
}

type RepoManage interface {
	Delete(param DeleteParam)
	Download(input string, output string)
	Upload(meta ComponentMeta, input string)
	List() []ComponentView
}
