package lg

import (
	"diana/config"
	"os"

	"github.com/rifflock/lfshook"
	"github.com/sirupsen/logrus"
	prefixed "github.com/x-cray/logrus-prefixed-formatter"
)

var log = logrus.New()
var stdFormatter *prefixed.TextFormatter  // 控制台日志输出格式
var fileFormatter *prefixed.TextFormatter // 文件日志输出格式

func Init() {
	c := config.GetConfig()

	timestampFormat := c.GetString("log.timestampFormat")

	stdFormatter = &prefixed.TextFormatter{
		FullTimestamp:   true,
		TimestampFormat: timestampFormat,
		ForceFormatting: true,
		ForceColors:     true,
		DisableColors:   false,
	}
	log.SetOutput(os.Stdout)
	log.SetFormatter(stdFormatter)

	fileFormatter = &prefixed.TextFormatter{
		FullTimestamp:   true,
		TimestampFormat: timestampFormat,
		ForceFormatting: true,
		ForceColors:     false,
		DisableColors:   true,
	}
	log.SetFormatter(fileFormatter)

	logLevel := c.GetString("log.level")
	switch logLevel {
	case "debug":
		// 设置日志级别为 debug 以上
		log.SetLevel(logrus.DebugLevel)
	case "info":
		// 设置日志级别为 info 以上
		log.SetLevel(logrus.InfoLevel)
	case "warn":
		// 设置日志级别为 warn 以上
		log.SetLevel(logrus.WarnLevel)
	case "error":
		// 设置日志级别为 error 以上
		log.SetLevel(logrus.ErrorLevel)
	case "fatal":
		// 设置日志级别为 fatal 以上
		log.SetLevel(logrus.FatalLevel)
	}

	logName := c.GetString("log.file.path")
	if logName != "" {
		pathMap := lfshook.PathMap{
			logrus.DebugLevel: logName,
			logrus.WarnLevel:  logName,
			logrus.InfoLevel:  logName,
			logrus.ErrorLevel: logName,
			logrus.FatalLevel: logName,
		}

		log.Hooks.Add(lfshook.NewHook(pathMap, &logrus.TextFormatter{}))
	}
}

// GetLog 获取日志对象
func GetLog() *logrus.Logger {
	return log
}
