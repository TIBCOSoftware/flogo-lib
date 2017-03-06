package logger

import (
	"fmt"
)

func Debug(args ...interface{}) {
	GetDefaultLogger().Debug(args...)
}

func Debugf(format string, args ...interface{}) {
	GetDefaultLogger().Debugf(format, args...)
}

func Info(args ...interface{}) {
	GetDefaultLogger().Info(args...)
}

func Infof(format string, args ...interface{}) {
	GetDefaultLogger().Infof(format, args...)
}

func Warn(args ...interface{}) {
	GetDefaultLogger().Warn(args...)
}

func Warnf(format string, args ...interface{}) {
	GetDefaultLogger().Warnf(format, args...)
}

func Error(args ...interface{}) {
	GetDefaultLogger().Error(args...)
}

func Errorf(format string, args ...interface{}) {
	GetDefaultLogger().Errorf(format, args...)
}

func SetLogLevel(level Level) {
	GetDefaultLogger().SetLogLevel(level)
}

func GetDefaultLogger() Logger {
	defLogger, err := GetLogger("default")
	if defLogger == nil {
		errorMsg := fmt.Sprintf("Engine: Error Getting Default Logger null")
		panic(errorMsg)
	}
	if err != nil {
		errorMsg := fmt.Sprintf("Engine: Error Getting Default Logger '%s'", err.Error())
		panic(errorMsg)
	}
	return defLogger
}
