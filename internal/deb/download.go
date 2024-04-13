package deb

import (
	"bufio"
	"fmt"
	"os"
	"strings"

	"github.com/baojingh/prctl/internal/logger"
	"github.com/baojingh/prctl/pkg/files"
	"github.com/baojingh/prctl/pkg/shell"
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
	doDownload(res)
	files.MoveFilesBatch("/var/cache/apt/archives/", output, ".deb")
	log.Info("Deb components are downloaded success.")
}

func doDownload(nameList string) {
	// component name list must be seperated and then composed by append.
	params := []string{"install", "--no-install-recommends", "-y", "--download-only"}
	params = append(params, strings.Fields(nameList)...)

	out, err := shell.DoShellCmd("apt-get", params...)
	if err != nil {
		log.Errorf("Failed to download %s, err: %s, out: %s", nameList, err, out)
		return
	}
	log.Infof("Download %s success.", nameList)
}
