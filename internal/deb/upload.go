package deb

import (
	"bytes"
	"net/http"
)

type DebComponentMeta struct {
	Distribution string
	Component    string
	Architech    string
}

// input: /xx/xx/xx/, check is it exists
func UploadDeb(meta DebComponentMeta, input string) {
	arch := meta.Architech
	dis := meta.Distribution
	com := meta.Component

	req, _ := http.NewRequest(http.MethodPut, "fullURL", bytes.NewReader(nil))
	req.Header.Set("Content-Type", "application/json")
	client := &http.Client{}
	client.Do(nil)

	log.Info("Deb components are uploaded success.", arch, dis, com)
}

func doUpload(name string) {

}
