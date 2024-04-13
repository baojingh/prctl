package deb

import (
	"fmt"
	"io"
	"net/http"
)

type DeleteParam struct {
	All     bool `json:"all" default:"false"`
	Name    string
	Version string
}

func DeleteDeb(param DeleteParam) {
	if param.All {
		deleteAll()
	} else if param.Version == "" {
		deleteByName()
	} else if param.Name != "" {
		deleteByVersion()
	}
}

func deleteByName() {

}

func deleteByVersion() {

}

func deleteAll() {
	// 替换为你的资源URL
	credPath := "/root/.prctl/.config"
	cred := ReadCred(credPath)
	// 创建一个HTTP DELETE请求
	req, err := http.NewRequest("DELETE", cred.Url, nil)
	if err != nil {
		log.Fatal(err)
	}
	req.SetBasicAuth(cred.Username, cred.Password)
	// 创建一个HTTP客户端
	client := &http.Client{}

	// 发送请求
	resp, err := client.Do(req)
	if err != nil {
		log.Fatal(err)
	}
	defer resp.Body.Close()

	// 读取响应体
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		log.Fatal(err)
	}

	// 打印响应状态码和响应体
	fmt.Printf("Response Status: %s\n", resp.Status)
	fmt.Printf("Response Body: %s\n", body)
}
