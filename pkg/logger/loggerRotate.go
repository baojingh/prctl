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

var logRotate = logrus.New()

type LogConfig struct {
	FileName string `json:"fileName" default:"log"`
	LogPath  string `json:"logPath" default:"./"`
	LogCount int    `json:"logCount" default:"1"`
	LogTime  int    `json:"logTime" default:"24"`
	LogLevel string `json:"logLevel" default:"info"`
}

// init logs
func configLogRotate(config LogConfig) {

	writer, err := rotate.New(
		filepath.Join(config.LogPath, fmt.Sprintf("%s-%s", config.FileName, "%Y%m%d")),
		rotate.WithLinkName(filepath.Join(config.LogPath, config.FileName)),
		rotate.WithRotationCount(uint(config.LogCount)),
		rotate.WithRotationTime(time.Hour*time.Duration(config.LogTime)),
	)
	if err == nil {
		logRotate.SetOutput(io.MultiWriter(writer, os.Stdout))
	} else {
		logRotate.Fatal("Config logger failure, ", err)
	}

	logRotate.SetFormatter(&logrus.TextFormatter{
		TimestampFormat:           "2006-01-02 15:04:05",
		ForceColors:               true,
		EnvironmentOverrideColors: true,
		FullTimestamp:             true,
		DisableLevelTruncation:    true,
	})
	logRotate.SetLevel(logrus.InfoLevel)
}

func NewRotate(config LogConfig) *logrus.Logger {
	configLogRotate(config)
	return logRotate
}
