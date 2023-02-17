package logger

import (
	"fmt"
	"github.com/sirupsen/logrus"
	"path"
	"runtime"
	"testing"
)

func Test_Logrus(t *testing.T) {
	logrus.SetReportCaller(true)
	logrus.SetFormatter(&MyFormatter{
		TextFormatter: &logrus.TextFormatter{
			CallerPrettyfier: func(f *runtime.Frame) (string, string) {
				_, filename := path.Split(f.File)
				filename = fmt.Sprintf("%s:%d", filename, f.Line)
				return "", filename
			},
		},
		CallerLevels: map[logrus.Level]bool{
			logrus.ErrorLevel: true,
			logrus.DebugLevel: true,
		},
	})

	logrus.Info("info level message")
	logrus.Error("error level message")
	logrus.Debug("debug level message")
	logrus.Debug("debug level message")
}
