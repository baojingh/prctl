package logger

import (
	"github.com/baojingh/prctl/pkg/logger"
	"github.com/sirupsen/logrus"
)

func New() *logrus.Logger {
	config := logger.LogConfig{
		FileName: "prctl.log",
	}
	log := logger.New(config)
	return log
}
