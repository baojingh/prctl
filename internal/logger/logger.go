package logger

import (
	"os/user"

	"github.com/baojingh/prctl/pkg/logger"
	"github.com/sirupsen/logrus"
)

func New() *logrus.Logger {
	currentUser, err := user.Current()
	if err != nil {
		return nil
	}
	config := logger.LogConfig{
		FileName: "prctl.log",
		LogPath:  currentUser.HomeDir,
	}
	log := logger.New(config)
	return log
}
