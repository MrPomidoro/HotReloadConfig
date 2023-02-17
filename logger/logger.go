package logger

import (
	"fmt"
	"github.com/sirupsen/logrus"
	"os"
	"path"
	"runtime"
)

type MyFormatter struct {
	*logrus.TextFormatter
	CallerLevels map[logrus.Level]bool
}

func New() *logrus.Logger {
	log := logrus.New()
	log.SetOutput(os.Stdout)
	log.SetLevel(logrus.DebugLevel)
	log.SetReportCaller(true)
	log.SetFormatter(&logrus.TextFormatter{
		CallerPrettyfier: func(f *runtime.Frame) (string, string) {
			_, filename := path.Split(f.File)
			filename = fmt.Sprintf(" %s:%d", filename, f.Line)
			return "", filename
		},
		FullTimestamp: true,
	})

	return log
}
