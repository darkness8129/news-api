package logging

import (
	"fmt"
	"os"

	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

var _ Logger = (*zapLogger)(nil)

type zapLogger struct {
	logger *zap.SugaredLogger
}

func NewZapLogger() (*zapLogger, error) {
	config := zap.Config{
		Encoding:         "console",
		Level:            zap.NewAtomicLevelAt(zap.DebugLevel),
		OutputPaths:      []string{"stderr"},
		ErrorOutputPaths: []string{"stderr"},
		EncoderConfig: zapcore.EncoderConfig{
			EncodeLevel:    zapcore.CapitalColorLevelEncoder,
			EncodeDuration: zapcore.SecondsDurationEncoder,
			EncodeName:     zapcore.FullNameEncoder,
			EncodeCaller:   zapcore.ShortCallerEncoder,
			EncodeTime:     zapcore.ISO8601TimeEncoder,
			LevelKey:       "severity",
			CallerKey:      "caller",
			TimeKey:        "timestamp",
			NameKey:        "name",
			MessageKey:     "message",
			LineEnding:     "\n",
		},
	}

	logger, err := config.Build()
	if err != nil {
		return nil, fmt.Errorf("failed to construct logger: %w", err)
	}

	return &zapLogger{logger.Sugar()}, nil
}

func (l *zapLogger) Named(name string) Logger {
	return &zapLogger{l.logger.Named(name)}
}

func (l *zapLogger) Debug(message string, args ...interface{}) {
	l.logger.Debugw(message, args...)
}

func (l *zapLogger) Info(message string, args ...interface{}) {
	l.logger.Infow(message, args...)
}

func (l *zapLogger) Error(message string, args ...interface{}) {
	l.logger.Errorw(message, args...)
}

func (l *zapLogger) Fatal(message string, args ...interface{}) {
	l.logger.Fatalw(message, args...)
	os.Exit(1)
}
