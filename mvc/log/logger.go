package log

import (
	"github.com/sirupsen/logrus"
	"os"
	"strings"
)

var (
	Logger   *logrus.Logger
	loglevel = "info"
)

func init() {
	level, err := logrus.ParseLevel(loglevel)
	if err != nil {
		level = logrus.DebugLevel
	}

	Logger = &logrus.Logger{
		Level: level,
		Out:   os.Stdout,
	}

	Logger.Formatter = &logrus.JSONFormatter{}

}

func Debug(msg string, tags ...string) {
	if Logger.Level < logrus.DebugLevel {
		return
	}
	Logger.WithFields(parseFields(tags...)).Debug(msg)
}

func Info(msg string, tags ...string) {
	if Logger.Level < logrus.InfoLevel {
		return
	}
	Logger.WithFields(parseFields(tags...)).Info(msg)
}

func parseFields(tags ...string) logrus.Fields {
	result := make(logrus.Fields, len(tags))
	for _,tag := range tags{
		els:=strings.Split(tag,":")
		result[strings.TrimSpace(els[0])]=strings.TrimSpace(els[1])
	}
	return result
}