package prhttp

import (
	"net/http"
	"sync"
	"time"
)

var once sync.Once
var client *http.Client

// Note: timeout when upload files with large size.
func init() {
	once.Do(func() {
		transport := &http.Transport{
			MaxIdleConnsPerHost: 60,
			IdleConnTimeout:     60,
		}

		client = &http.Client{
			Timeout:   time.Second * 60,
			Transport: transport,
		}
	})
}

func GetHttpClient() *http.Client {
	return client
}

func DoHttpRequest(req *http.Request) (*http.Response, error) {
	client := GetHttpClient()
	resp, err := client.Do(req)
	if err != nil {
		return nil, err
	}
	return resp, nil
}
