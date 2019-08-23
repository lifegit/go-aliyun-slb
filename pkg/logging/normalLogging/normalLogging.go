package normalLogging

import (
	"github.com/sirupsen/logrus"
	"go-aliyun-slb/pkg/logging"
)

var Logger *logrus.Logger

func Setup(dir string, exitHandler func()) {
	Logger = logging.NewLogger(dir, 3, &logrus.JSONFormatter{}, exitHandler)
}
