package logger

import (
	"github.com/sirupsen/logrus"
)

var log = logrus.New()

// init logs
func configLog(config LogConfig) {
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
