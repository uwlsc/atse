package lib

import (
	"io"
	"os"

	"go.uber.org/fx"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

var globalLog *Logger
var zapLogger *zap.Logger

// Logger structure
type Logger struct {
	*zap.SugaredLogger
}

// FxLogger logger for go-fx [subbed from main logger]
type FxLogger struct {
	*Logger
}

// GinLogger logger for gin framework [subbed from main logger]
type GinLogger struct {
	*Logger
}

// GetLogger gets the global instance of the logger
func GetLogger() Logger {
	if globalLog != nil {
		return *globalLog
	}
	globalLog := newLogger()
	return *globalLog
}

// newLogger sets up logger the main logger
func newLogger() *Logger {

	env := os.Getenv("ENVIRONMENT")
	config := zap.NewDevelopmentConfig()

	if env == "local" {
		config.EncoderConfig.EncodeLevel = zapcore.CapitalColorLevelEncoder
	} else {
		config.Level.SetLevel(zap.PanicLevel)
	}

	zapLogger, _ = config.Build()

	globalLog := zapLogger.Sugar()

	return &Logger{
		SugaredLogger: globalLog,
	}
}

func newSugaredLogger(logger *zap.Logger) *Logger {
	return &Logger{
		SugaredLogger: logger.Sugar(),
	}
}

// GetFxLogger gets logger for go-fx
func (l *Logger) GetFxLogger() fx.Printer {
	logger := zapLogger.WithOptions(
		zap.WithCaller(false),
	)
	return FxLogger{
		Logger: newSugaredLogger(logger),
	}
}

// GetGinLogger gets logger for gin framework debugging
func (l *Logger) GetGinLogger() io.Writer {
	logger := zapLogger.WithOptions(
		zap.WithCaller(false),
	)
	return GinLogger{
		Logger: newSugaredLogger(logger),
	}
}

// Printf prints go-fx logs
func (l FxLogger) Printf(str string, args ...interface{}) {
	l.Infof(str, args)
}

// Writer interface implementation for gin-framework
func (l GinLogger) Write(p []byte) (n int, err error) {
	l.Info(string(p))
	return len(p), nil
}
