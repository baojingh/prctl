package deb

// import (
// 	"bufio"
// 	"fmt"
// 	"os"
// 	"sync"

// 	"github.com/baojingh/prctl/internal/grpool"
// 	"github.com/baojingh/prctl/internal/logger"
// 	"github.com/baojingh/prctl/internal/utils/shell"
// )

// var log = logger.New()

// func checkDebEnv() {
// 	fmt.Println("hello")
// }

// func prepareDebEnv() {
// 	_, err := shell.DoShellCmd("apt-get", "update")
// 	if err != nil {
// 		return
// 	}

// }

// // input: /xx/xx/xx/ss.txt, check is it exists
// // output aa/ss/ created if not exist
// func DownloadDeb(input string, output string) {
// 	// prepareDebEnv()

// 	file, err := os.Open(input)
// 	if err != nil {
// 		log.Errorf("Cannot open file %s", input)
// 		return
// 	}
// 	defer file.Close()

// 	var wg sync.WaitGroup

// 	scanner := bufio.NewScanner(file)
// 	for scanner.Scan() {
// 		name := scanner.Text()
// 		wg.Add(1)
// 		f := func() {
// 			doDownload(name, &wg)
// 		}
// 		grpool.SubmitTask(f)
// 	}
// 	if err := scanner.Err(); err != nil {
// 		return
// 	}
// 	wg.Wait()
// 	log.Info("Deb components are downloaded success.")
// }

// func doDownload(name string, wg *sync.WaitGroup) {
// 	defer wg.Done()
// 	params := []string{"install", "--no-install-recommends", "-y", "--download-only", name}
// 	out, err := shell.DoShellCmd("apt-get", params...)
// 	if err != nil {
// 		log.Errorf("Failed to download %s, err: %s, out: %s", name, err, out)
// 		return
// 	}
// 	log.Infof("Download %s success.", name)
// }
