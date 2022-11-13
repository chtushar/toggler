package logger

import (
	"log"

	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

type Config struct {
	Production bool
}

// New creates a new instance of the logger based on the config
func New(cfg *Config) *zap.Logger {
	if cfg.Production {
		logger, err := zap.NewProduction()
		if err != nil {
			log.Panicf("Failed to initialize production logger: %v", err)
		}

		return logger
	}

	devConfig := zap.NewDevelopmentConfig()
	devConfig.EncoderConfig.EncodeLevel = zapcore.CapitalColorLevelEncoder

	logger, err := devConfig.Build()
	if err != nil {
		log.Panicf("Failed to initialize development logger: %v", err)
	}

	return logger
}
