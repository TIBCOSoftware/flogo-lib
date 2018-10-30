package logger

import (
	"fmt"
)

func Debug(args ...interface{}) {
	defaultLogger.Debug(args...)
}

func Debugf(format string, args ...interface{}) {
	defaultLogger.Debugf(format, args...)
}

func Info(args ...interface{}) {
	defaultLogger.Info(args...)
}

func Infof(format string, args ...interface{}) {
	defaultLogger.Infof(format, args...)
}

func Warn(args ...interface{}) {
	defaultLogger.Warn(args...)
}

func Warnf(format string, args ...interface{}) {
	defaultLogger.Warnf(format, args...)
}

func Error(args ...interface{}) {
	defaultLogger.Error(args...)
}

func Errorf(format string, args ...interface{}) {
	defaultLogger.Errorf(format, args...)
}

func SetLogLevel(level Level) {
	defaultLogger.SetLogLevel(level)
}

func GetLogLevel() Level {
	return defaultLogger.GetLogLevel()
}

var defaultLoggerName = "flogo"
var defaultLogLevel = "INFO"
var defaultLogger Logger

func SetDefaultLogger(name string) {
	defaultLoggerName = name
}

func GetDefaultLogger() Logger {
	return defaultLogger
}

func getDefaultLogger() Logger {
	defLogger := GetLogger(defaultLoggerName)
	if defLogger == nil {
		errorMsg := fmt.Sprintf("error getting default logger '%s'", defaultLoggerName)
		panic(errorMsg)
	}
	return defLogger
}


