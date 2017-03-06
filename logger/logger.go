package logger

import (
	"errors"
	"fmt"
)

type Logger interface {
	Debug(args ...interface{})
	Debugf(format string, args ...interface{})
	Info(args ...interface{})
	Infof(format string, args ...interface{})
	Warn(args ...interface{})
	Warnf(format string, args ...interface{})
	Error(args ...interface{})
	Errorf(format string, args ...interface{})
	SetLogLevel(Level)
}

type LoggerFactory interface {
	GetLogger(name string) (Logger, error)
}

type Level int

const (
	DebugLevel Level = iota
	InfoLevel
	WarnLevel
	ErrorLevel
)

var levelNames = initLevelNames()

func initLevelNames() map[string]Level {
	newLevelNames := make(map[string]Level, 4)
	newLevelNames["DEBUG"] = DebugLevel
	newLevelNames["INFO"] = InfoLevel
	newLevelNames["WARN"] = WarnLevel
	newLevelNames["ERROR"] = ErrorLevel
	return newLevelNames
}

var logFactory LoggerFactory

func RegisterLoggerFactory(factory LoggerFactory) {
	logFactory = factory
}

// GetLogger returns the Logger using the logFactory registered.
func GetLogger(name string) (Logger, error) {
	if logFactory == nil {
		return nil, errors.New("No logger factory found.")
	}
	return logFactory.GetLogger(name)
}

func GetLevelForName(name string) (Level, error) {
	levelForName, ok := levelNames[name]
	if !ok {
		return 0, fmt.Errorf("Unsupported Log Level '%s'", name)
	}
	return levelForName, nil
}
