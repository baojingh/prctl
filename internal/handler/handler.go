package handler

type ClientOperations interface {
	Delete(param string)
	Download(param string)
	Upload(param string)
}
