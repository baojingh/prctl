package deb

import (
	"net/http"
	"sync"
)

var once sync.Once
var client *http.Client

func init() {
	once.Do(func() {
		client = &http.Client{}
	})
}

func GetHttpClient() *http.Client {
	return client
}

// import (
// 	"bytes"
// 	"fmt"
// 	"io"
// 	"mime/multipart"
// 	"os"
// 	"path/filepath"
// )

// func doRequst(meta DebComponentMeta, cred CredentialInfo, path string, name string) {
// 	absFilePath := filepath.Join(path, name)
// 	file, _ := os.Open(absFilePath)
// 	defer file.Close()
// 	body := new(bytes.Buffer)
// 	writer := multipart.NewWriter(body)
// 	part, _ := writer.CreateFormFile("file", file.Name())
// 	io.Copy(part, file)
// 	writer.Close()

// 	uploadUrl := fmt.Sprintf("%s/%s;deb.distribution=%s;deb.component=%s;deb.architecture=%s",
// 		cred.Url, name, meta.Distribution, meta.Component, meta.Architech)
// 	log.Infoln(cred.Url, cred.Username, cred.Password)
// 	log.Infoln(path, name, uploadUrl)

// 	// curl -u${USER}:${TOKEN} \
// 	//      -XPUT  \
// 	//     "${URL}/${file_name};deb.distribution=${DISTRIBUTION};deb.component=${COMPONENT};deb.architecture=${ARCH}" \
// 	//     -T "${file}"
// 	// req, _ := http.NewRequest("PUT", uploadUrl, body)

// 	// req.SetBasicAuth(cred.Username, cred.Password)
// 	// req.Header.Set("Content-Type", writer.FormDataContentType())
// 	// client := &http.Client{}
// 	// resp, err := client.Do(req)
// 	// if err != nil {
// 	// 	log.Errorf("Fail to send request, %s", err)
// 	// }
// 	// defer resp.Body.Close()
// 	// // 打印响应
// 	// log.Infof("Response Status: %s, body: %s\n", resp.Status, resp.Body)
// }

// func doResponse() {

// }
