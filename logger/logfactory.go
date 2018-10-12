package logger

import (
	"fmt"
	"sync"

	"github.com/sirupsen/logrus"
	"github.com/TIBCOSoftware/flogo-lib/config"
	"strings"
)

var loggerMap = make(map[string]Logger)
var overrideLogLevelMap = make(map[string]Level)
var mutex = &sync.RWMutex{}
var logLevel = InfoLevel
var logFormat = "TEXT"

type DefaultLoggerFactory struct {
}

func init() {

	RegisterLoggerFactory(&DefaultLoggerFactory{})

	logLevelName := config.GetLogLevel()
	// Get log level for name
	getLogLevel, err := GetLevelForName(logLevelName)
	if err != nil {
		println("Unsupported Log Level - [" + logLevelName + "]. Set to Log Level - [" + defaultLogLevel + "]")
	} else {
		logLevel = getLogLevel
	}

	// Gather overridden log levels
	overrideValues := config.GetContribLogLevelOverride()
	if overrideValues != "" {
		for _, pair := range strings.Split(overrideValues, ",") {
			kv := strings.Split(pair, "=")
			if len(kv) == 2 && kv[0] != "" &&  kv[1] != "" {
				ll, err := GetLevelForName(kv[1])
				if err != nil {
					println("'%s' is not valid override value for '%s'. %s", pair, config.ENV_LOG_LEVEL_OVERRIDE_KEY, err.Error())
				} else {
					overrideLogLevelMap[kv[0]] = ll
				}
			} else {
				println("'%s' is not valid override value for '%s'. It must be in PropName=PropValue format.", pair, config.ENV_LOG_LEVEL_OVERRIDE_KEY)
			}
		}
	}

	logFormat = config.GetLogFormat()
}

type DefaultLogger struct {
	loggerName string
	loggerImpl *logrus.Logger
}




type LogFormatter struct {
	loggerName string
}

func (f *LogFormatter) Format(entry *logrus.Entry) ([]byte, error) {
	logEntry := fmt.Sprintf("%s %-6s [%s] - %s\n", entry.Time.Format(config.GetLogDateTimeFormat()), getLevel(entry.Level), f.loggerName, entry.Message)
	return []byte(logEntry), nil
}



func getLevel(level logrus.Level) string {
	switch level {
	case logrus.DebugLevel:
		return "DEBUG"
	case logrus.InfoLevel:
		return "INFO"
	case logrus.ErrorLevel:
		return "ERROR"
	case logrus.WarnLevel:
		return "WARN"
	case logrus.PanicLevel:
		return "PANIC"
	case logrus.FatalLevel:
		return "FATAL"
	}

	return "UNKNOWN"
}

// Debug logs message at Debug level.
func (l *DefaultLogger) Debug(args ...interface{}) {
	l.loggerImpl.Debug(args...)
}

// DebugEnabled checks if Debug level is enabled.
func (l *DefaultLogger) DebugEnabled() bool {
	return l.loggerImpl.Level >= logrus.DebugLevel
}

// Info logs message at Info level.
func (l *DefaultLogger) Info(args ...interface{}) {
	l.loggerImpl.Info(args...)
}

// InfoEnabled checks if Info level is enabled.
func (l *DefaultLogger) InfoEnabled() bool {
	return l.loggerImpl.Level >= logrus.InfoLevel
}

// Warn logs message at Warning level.
func (l *DefaultLogger) Warn(args ...interface{}) {
	l.loggerImpl.Warn(args...)
}

// WarnEnabled checks if Warning level is enabled.
func (l *DefaultLogger) WarnEnabled() bool {
	return l.loggerImpl.Level >= logrus.WarnLevel
}

// Error logs message at Error level.
func (l *DefaultLogger) Error(args ...interface{}) {
	l.loggerImpl.Error(args...)
}

// ErrorEnabled checks if Error level is enabled.
func (l *DefaultLogger) ErrorEnabled() bool {
	return l.loggerImpl.Level >= logrus.ErrorLevel
}

// Debug logs message at Debug level.
func (l *DefaultLogger) Debugf(format string, args ...interface{}) {
	l.loggerImpl.Debugf(format, args...)
}

// Info logs message at Info level.
func (l *DefaultLogger) Infof(format string, args ...interface{}) {
	l.loggerImpl.Infof(format, args...)
}

// Warn logs message at Warning level.
func (l *DefaultLogger) Warnf(format string, args ...interface{}) {
	l.loggerImpl.Warnf(format, args...)
}

// Error logs message at Error level.
func (l *DefaultLogger) Errorf(format string, args ...interface{}) {
	l.loggerImpl.Errorf(format, args...)
}

//SetLog Level
func (l *DefaultLogger) SetLogLevel(logLevel Level) {
	switch logLevel {
	case DebugLevel:
		l.loggerImpl.SetLevel(logrus.DebugLevel)
	case InfoLevel:
		l.loggerImpl.SetLevel(logrus.InfoLevel)
	case ErrorLevel:
		l.loggerImpl.SetLevel(logrus.ErrorLevel)
	case WarnLevel:
		l.loggerImpl.SetLevel(logrus.WarnLevel)
	default:
		l.loggerImpl.SetLevel(logrus.ErrorLevel)
	}
}

func (l *DefaultLogger) GetLogLevel() Level {
	levelStr := getLevel(l.loggerImpl.Level)
	level, _ := GetLevelForName(levelStr)
	return level
}

func getLogLevel(loggerName string) Level {
	if len(overrideLogLevelMap) > 0 {
		// Find overridden log level
		for name, loglevel := range overrideLogLevelMap {
			// Look for logger which matches given name
			if strings.Contains(loggerName, name) {
				return loglevel
			}
		}
	}
	// Return engine log level
	return logLevel
}

func (logfactory *DefaultLoggerFactory) GetLogger(name string) Logger {
	mutex.RLock()
	l := loggerMap[name]
	mutex.RUnlock()
	if l == nil {
		logImpl := logrus.New()

		if logFormat == "JSON" {
			logImpl.Formatter = &logrus.JSONFormatter{}
		} else {
			logImpl.Formatter = &LogFormatter{
				loggerName: name,
			}
		}

		l = &DefaultLogger{
			loggerName: name,
			loggerImpl: logImpl,
		}

		l.SetLogLevel(getLogLevel(name))

		mutex.Lock()
		loggerMap[name] = l
		mutex.Unlock()
	}
	return l
}
