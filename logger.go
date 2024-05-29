package utils

import (
	"github.com/sirupsen/logrus"
	"go.opentelemetry.io/otel/log"
)

type Logger interface {
	LogError(err error,opts ...OptionLog)
	LogWarn(m string,opts ...OptionLog)
	LogInfo(m string,opts ...OptionLog)
}

var AllawordLevels = []logrus.Level{
	logrus.Level(log.SeverityDebug),
	logrus.Level(log.SeverityTrace),
	logrus.Level(log.SeverityInfo),
	logrus.Level(log.SeverityWarn),
	logrus.Level(log.SeverityError),
	logrus.Level(log.SeverityFatal),
}
