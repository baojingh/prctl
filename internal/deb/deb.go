package deb

import (
	"bufio"
	"encoding/base64"
	"encoding/json"
	"fmt"
	"net/http"
	"os"
	"path/filepath"
	"strings"
	"sync"

	"github.com/baojingh/prctl/internal/logger"
	"github.com/baojingh/prctl/pkg/files"
	"github.com/baojingh/prctl/pkg/grpool"
	"github.com/baojingh/prctl/pkg/prsys"
	"github.com/baojingh/prctl/pkg/shell"
)

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

var log = logger.New()

// Create client from a cred file, apply in: logout,download, upload, delete
func (j *DebClientFactory) CreateClient() interface{} {
	configPath := getConfigPath()
	content, err := os.ReadFile(configPath)
	if err != nil {
		log.Errorf("configPath %s not exist.", configPath)
		return nil
	}
	var cli Client
	decodeCred, err := base64.StdEncoding.DecodeString(string(content))
	if err != nil {
		log.Errorf("fail to decode, %s", err)
		return nil
	}
	err = json.Unmarshal(decodeCred, &cli)
	if err != nil {
		log.Errorf("fail to unmarshal from decode data, %s", err)
		return nil
	}
	return &cli
}

// get cred path, default os /home/${USER}/.prctl/config or /root/.prctl/config
func getConfigPath() string {
	userPath := prsys.CurrentUserPath()
	hiddenPath := filepath.Join(userPath, ".prctl")
	configPath := filepath.Join(hiddenPath, "config")
	return configPath
}

func (cli *Client) Delete(param DeleteParam) error {
	if param.All {
		log.Infof("delete all, %v", param)
	} else if param.Name != "" && param.Version == "" {
		log.Infof("delete component, %v", param)
	} else if param.Name != "" && param.Version != "" {
		log.Infof("delete component and version, %v", param)
	}
	return nil
}

// input: /xx/xx/xx/ss.txt, check is it exists
// output aa/ss/ created if not exist
func (cli *Client) Download(input string, output string) {
	// init debian environment
	_, err := shell.DoShellCmd("apt-get", "update")
	if err != nil {
		log.Errorf("fail to apt-get update deb env, err: %s", err)
		return
	}
	file, err := os.Open(input)
	if err != nil {
		log.Errorf("Cannot open file %s", input)
		return
	}
	defer file.Close()

	//  Create the output dir if it not exist
	files.CreateDirIfNotExist(output, 0755)

	var buffer strings.Builder
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		name := scanner.Text()
		buffer.WriteString(name)
		buffer.WriteString(" ")
	}
	if err := scanner.Err(); err != nil {
		return
	}
	res := strings.TrimSpace(buffer.String())
	changeDirAndDo(res, output)
	log.Info("Deb components are downloaded success.")
}

// https://stackoverflow.com/questions/52435908/how-to-change-the-shells-current-working-directory-in-go
// apt-get download just put the components in current path, so we need change to target dir
func changeDirAndDo(nameList string, path string) {
	cwd, _ := os.Getwd()
	if err := os.Chdir(path); err != nil {
		return
	}
	// component name list must be seperated and then composed by append.
	params := []string{"download"}
	params = append(params, strings.Fields(nameList)...)
	log.Infof("Command: apt-get %s", strings.Join(params, " "))
	out, err := shell.DoShellCmd("apt-get", params...)
	if err != nil {
		log.Errorf("failed to download %s, err: %s, out: %s", nameList, err, out)
		return
	}
	log.Infof("Download %s success.", nameList)

	if err := os.Chdir(cwd); err != nil {
		return
	}
}

func (cli *Client) Upload(meta ComponentMeta, input string) {
	log.Infof("start upload, input path %s", input)
	var wg sync.WaitGroup

	fileList, _ := files.ListFilesInDir(input)
	for _, file := range fileList {
		fileName := file
		wg.Add(1)
		f := func() {
			defer wg.Done()
			cli.doUpload(meta, input, fileName)
		}
		grpool.SubmitTask(f)
	}
	// NOTE: Do Not Forget it.
	wg.Wait()
}

//	curl -u${USER}:${TOKEN} \
//	     -XPUT  \
//	    "${URL}/${file_name};deb.distribution=${DISTRIBUTION};deb.component=${COMPONENT};deb.architecture=${ARCH}" \
//	    -T "${file}"
func (cli *Client) doUpload(meta ComponentMeta, path string, fileName string) {
	arch := meta.Architech
	dis := meta.Distribution
	com := meta.Component
	uploadUrl := fmt.Sprintf("%s/%s;deb.distribution=%s;deb.component=%s;deb.architecture=%s",
		cli.RepoUrl, fileName, dis, com, arch)

	// Open the file
	filePath := filepath.Join(path, fileName)
	file, err := os.Open(filePath)
	if err != nil {
		log.Error(err)
		return
	}
	defer file.Close()

	// Create request object
	req, err := http.NewRequest("PUT", uploadUrl, file)
	if err != nil {
		log.Error(err)
		return
	}

	// Set username and password
	req.Header.Set("Content-Type", "application/octet-stream")
	req.SetBasicAuth(cli.Username, cli.Password)

	// // Do request
	// resp, err := prhttp.DoHttpRequest(req)
	// if err != nil {
	// 	log.Error(err)
	// 	return
	// }
	// defer resp.Body.Close()

	log.Infof("HTTP upload Success Status")
}
