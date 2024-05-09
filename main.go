package main

import (
	"os"

	"github.com/baojingh/prctl/cmd"
	"github.com/baojingh/prctl/internal/logger"
)

var log = logger.New()

func main() {
	err := cmd.RootCmd.Execute()
	if err != nil {
		log.Fatalf("Command failed, %s", err)
		os.Exit(1)
	}
	os.Exit(0)

}
