package logger

import (
	"fmt"
	"strings"

	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

var rootLogger Logger

func initRootLogger() {
	zl, lvl, _ := newZapLogger()
	rootLogger = &zapLoggerImpl{loggerLevel: lvl, mainLogger: zl.Sugar()}
}

type ZapLoggerFactory struct {
}

func (*ZapLoggerFactory) GetLogger(name string) Logger {
	mutex.RLock()
	l := loggerMap[name]
	mutex.RUnlock()
	
	if l == nil {
		var err error
		l, err = newZapChildLogger(rootLogger, name)
		if err != nil {
			return rootLogger
		}

		mutex.Lock()
		loggerMap[name] = l
		mutex.Unlock()
	}
	return l
}

type zapLoggerImpl struct {
	loggerLevel *zap.AtomicLevel
	mainLogger  *zap.SugaredLogger
}

func (l *zapLoggerImpl) SetLogLevel(level Level) {
	zapLevel := toZapLogLevel(level)
	l.loggerLevel.SetLevel(zapLevel)
}

func (l *zapLoggerImpl) GetLogLevel() Level {
	zapLevel := l.loggerLevel.Level()
	return fromZapLogLevel(zapLevel)
}

func (l *zapLoggerImpl) DebugEnabled() bool {
	return l.loggerLevel.Enabled(zapcore.DebugLevel)
}

func (l *zapLoggerImpl) Debug(args ...interface{}) {
	l.mainLogger.Debug(args...)
}

func (l *zapLoggerImpl) Info(args ...interface{}) {
	l.mainLogger.Info(args...)
}

func (l *zapLoggerImpl) Warn(args ...interface{}) {
	l.mainLogger.Warn(args...)
}

func (l *zapLoggerImpl) Error(args ...interface{}) {
	l.mainLogger.Error(args...)
}

func (l *zapLoggerImpl) Debugf(template string, args ...interface{}) {
	l.mainLogger.Debugf(template, args...)
}

func (l *zapLoggerImpl) Infof(template string, args ...interface{}) {
	l.mainLogger.Infof(template, args...)
}

func (l *zapLoggerImpl) Warnf(template string, args ...interface{}) {
	l.mainLogger.Warnf(template, args...)
}

func (l *zapLoggerImpl) Errorf(template string, args ...interface{}) {
	l.mainLogger.Errorf(template, args...)
}

func toZapLogLevel(level Level) zapcore.Level {
	switch level {
	case DebugLevel:
		return zapcore.DebugLevel
	case InfoLevel:
		return zapcore.InfoLevel
	case WarnLevel:
		return zapcore.WarnLevel
	case ErrorLevel:
		return zapcore.ErrorLevel
	}

	return zapcore.InfoLevel
}

func fromZapLogLevel(level zapcore.Level) Level {
	switch level {
	case zapcore.DebugLevel:
		return DebugLevel
	case zapcore.InfoLevel:
		return InfoLevel
	case zapcore.WarnLevel:
		return WarnLevel
	case zapcore.ErrorLevel:
		return ErrorLevel
	}

	return InfoLevel
}

func newZapLogger() (*zap.Logger, *zap.AtomicLevel, error) {
	cfg := zap.NewProductionConfig()
	cfg.DisableCaller = true

	eCfg := cfg.EncoderConfig
	eCfg.TimeKey = "timestamp"
	eCfg.EncodeTime = zapcore.ISO8601TimeEncoder

	if !strings.EqualFold(logFormat, "JSON") {
		eCfg.EncodeLevel = zapcore.CapitalLevelEncoder
		cfg.Encoding = "console"
		eCfg.EncodeName = nameEncoder
	}

	cfg.EncoderConfig = eCfg

	lvl := cfg.Level
	zl, err := cfg.Build(zap.AddCallerSkip(1))

	return zl, &lvl, err
}

func nameEncoder(loggerName string, enc zapcore.PrimitiveArrayEncoder) {
	enc.AppendString("[" + loggerName + "] -")
}

func newZapChildLogger(l Logger, name string) (Logger, error) {

	impl, ok := l.(*zapLoggerImpl)

	if ok {
		zapLogger := impl.mainLogger
		newZl := zapLogger.Named(name)

		return &zapLoggerImpl{loggerLevel: impl.loggerLevel, mainLogger: newZl}, nil
	} else {
		return nil, fmt.Errorf("unable to create child logger")
	}
}

func zapSync(l Logger) {
	impl, ok := l.(*zapLoggerImpl)

	if ok {
		impl.mainLogger.Sync()
	}
}
