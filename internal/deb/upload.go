package deb

type DebComponentMeta struct {
	distribution string
	component    string
	architech    string
}

// input: /xx/xx/xx/, check is it exists
func UploadDeb(meta DebComponentMeta, input string) {

	log.Info("Deb components are uploaded success.")
}

func doUpload(name string) {

}
