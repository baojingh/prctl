package handler

type RepoManage interface {
	Delete(param string)
	Download(param string)
	Upload(param string)
}
