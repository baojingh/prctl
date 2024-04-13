package logger

import (
	"fmt"
	"io"
	"os"
	"path/filepath"
	"time"

	rotate "github.com/lestrrat-go/file-rotatelogs"
	"github.com/sirupsen/logrus"
)

var log = logrus.New()

type LogConfig struct {
	FileName string `json:"fileName" default:"log"`
	LogPath  string `json:"logPath" default:"./"`
	LogCount int    `json:"logCount" default:"1"`
	LogTime  int    `json:"logTime" default:"24"`
	LogLevel string `json:"logLevel" default:"info"`
}

// init logs
func configLog(config LogConfig) {

	writer, err := rotate.New(
		filepath.Join(config.LogPath, fmt.Sprintf("%s-%s", config.FileName, "%Y%m%d")),
		rotate.WithLinkName(filepath.Join(config.LogPath, config.FileName)),
		rotate.WithRotationCount(uint(config.LogCount)),
		rotate.WithRotationTime(time.Hour*time.Duration(config.LogTime)),
	)
	if err == nil {
		log.SetOutput(io.MultiWriter(writer, os.Stdout))
	} else {
		log.Fatal("Config logger failure, ", err)
	}

	log.SetFormatter(&logrus.TextFormatter{
		TimestampFormat:           "2006-01-02 15:04:05",
		ForceColors:               true,
		EnvironmentOverrideColors: true,
		FullTimestamp:             true,
		DisableLevelTruncation:    true,
	})
	log.SetLevel(logrus.InfoLevel)
}

func New(config LogConfig) *logrus.Logger {
	configLog(config)
	return log
}
