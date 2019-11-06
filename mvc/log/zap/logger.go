package zap

import (
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

var (
	Logger *zap.Logger
)

func init() {
	loggerConfig := zap.Config{
		OutputPaths: []string{"stdout"},
		Encoding:    "json",
		Level:       zap.NewAtomicLevel(),
		EncoderConfig: zapcore.EncoderConfig{
			MessageKey:   "msg",
			LevelKey:     "level",
			TimeKey:      "time",
			EncodeTime:   zapcore.ISO8601TimeEncoder,
			EncodeLevel:  zapcore.LowercaseLevelEncoder,
			EncodeCaller: zapcore.ShortCallerEncoder,
		},
	}
	var err error
	Logger, err = loggerConfig.Build()
	if err != nil {
		panic(err)
	}
}

func Debug(msg string, tags ...zap.Field) {
	Logger.Debug(msg, tags...)
	Logger.Sync()
}

func Info(msg string, tags ...zap.Field) {
	Logger.Info(msg, tags...)
	Logger.Sync()
}
