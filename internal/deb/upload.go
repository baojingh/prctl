package deb

import (
	"github.com/baojingh/prctl/internal/utils/files"
)

type DebComponentMeta struct {
	Distribution string
	Component    string
	Architech    string
}

// input: /xx/xx/xx/, check is it exists
func UploadDeb(meta DebComponentMeta, input string) {
	// arch := meta.Architech
	// dis := meta.Distribution
	// com := meta.Component
	log.Infof("start upload, input path %s", input)
	// var wg sync.WaitGroup
	fileList, _ := files.ListFilesInDir(input)
	for _, file := range fileList {
		fileName := file
		doUpload(meta, input, fileName)
		// wg.Add(1)
		// f := func() {
		// 	defer wg.Done()
		// 	doUpload(input, fileName)
		// }
		// grpool.SubmitTask(f)
	}
	// NOTE: Do Not Forget it.
	// wg.Wait()
}

func doUpload(meta DebComponentMeta, path string, name string) {
	credPath := "/root/.prctl/.config"
	cred := ReadCred(credPath)
	doRequst(meta, cred, path, name)

	// file, _ := os.Open(absFilePath)
	// defer file.Close()
	// body := new(bytes.Buffer)
	// writer := multipart.NewWriter(body)
	// part, _ := writer.CreateFormFile("file", file.Name())
	// io.Copy(part, file)
	// writer.Close()

	// uploadUrl := fmt.Sprintf("%s/%s;deb.distribution=%s;deb.component=%s;deb.architecture=%s",
	// 	cred.Url, name, "dis", "com", "arch")

	// // curl -u${USER}:${TOKEN} \
	// //      -XPUT  \
	// //     "${URL}/${file_name};deb.distribution=${DISTRIBUTION};deb.component=${COMPONENT};deb.architecture=${ARCH}" \
	// //     -T "${file}"
	// req, _ := http.NewRequest("PUT", uploadUrl, body)

	// req.SetBasicAuth(cred.Username, cred.Password)
	// req.Header.Set("Content-Type", writer.FormDataContentType())
	// client := &http.Client{}
	// resp, _ := client.Do(req)
	// defer resp.Body.Close()
	// // 打印响应
	// log.Infof("Response Status: %s, body: %s\n", resp.Status, resp.Body)
}
