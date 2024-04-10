package deb

import (
	"bufio"
	"fmt"
	"os"

	"github.com/baojingh/prctl/internal/logger"
	"github.com/baojingh/prctl/internal/utils/shell"
)

var log = logger.New()

func checkDebEnv() {
	fmt.Println("hello")
}

func prepareDebEnv() {
	_, err := shell.DoShellCmd("apt-get", "update")
	if err != nil {
		return
	}

}

// input: /xx/xx/xx/ss.txt, check is it exists
// output aa/ss/ created if not exist
func DownloadDeb(input string, output string) {
	// prepareDebEnv()

	file, err := os.Open(input)
	if err != nil {
		log.Errorf("Cannot open file %s", input)
		return
	}
	defer file.Close()
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()
		doDownload(line)
	}
	if err := scanner.Err(); err != nil {
		return
	}
}

func doDownload(name string) (string, error) {
	// log.Infof("hello, %s", name)
	// return "", nil
	// param := fmt.Sprintf("install --no-install-recommends -y --download-only ")
	// log.Info(param)
	// out, err := shell.DoShellCmd("apt-get", param)
	params := []string{"install", "--no-install-recommends", "-y", "--download-only", name}
	out, err := shell.DoShellCmd("apt-get", params...)
	// out, err := shell.DoShellCmd("apt-get", "install --no-install-recommends -y --download-only", "gosu")
	log.Infof("out: %s, err: %s", out, err)

	return out, err

	// defer file.Close()
	// scanner := bufio.NewScanner(file)
	//
	//	for scanner.Scan() {
	//		line := scanner.Text()
	//		fields := strings.Split(line, "=")
	//		osMap[]
	//		if len(fields) >= 2 && fields[0] == "ID" {
	//			dis := strings.TrimSpace(fields[1])
	//			return dis, nil
	//		}
	//	}
	//
	//	if scanner.Err(); err != nil {
	//		return "", err
	//	}

}
