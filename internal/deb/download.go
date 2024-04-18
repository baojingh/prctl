package deb

import (
	"bufio"
	"os"
	"strings"

	"github.com/baojingh/prctl/pkg/files"
	"github.com/baojingh/prctl/pkg/shell"
)

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
