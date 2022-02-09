package log

import (
	"github.com/sirupsen/logrus"
	"os"
)

var (
	Log = logrus.New()
)

func init() {

	Log.SetReportCaller(true)
	Log.SetOutput(os.Stdout)
	Log.SetFormatter(&CustomFormatter{})
}
