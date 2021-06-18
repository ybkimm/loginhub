package main

import (
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

var logger = func() *zap.Logger {
	cfg := zap.NewDevelopmentConfig()
	cfg.EncoderConfig.EncodeLevel = zapcore.CapitalColorLevelEncoder
	logger, err := cfg.Build()
	if err != nil {
		panic(err)
	}
	return logger
}()
