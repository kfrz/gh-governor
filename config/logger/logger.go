package logger

import (
	"log"
	"os"
	"strings"

	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

// createLogger() creates a new zap logger from settings in the environment
// variables. If no settings are found, it defaults to INFO level logging.
// Default to production configuration unless APP_ENV=development
func createLogger() *zap.Logger {
	logLevel := os.Getenv("LOG_LEVEL")
	if logLevel == "" {
		logLevel = "info"
	}

	var level zapcore.Level
	if err := level.UnmarshalText([]byte(strings.ToLower(logLevel))); err != nil {
		log.Printf("Error setting log level: %v. Defaulting to INFO level.", err)
		level = zapcore.InfoLevel
	}

	developmentCfg := zap.Config{
		Encoding:         "console",
		Level:            zap.NewAtomicLevelAt(level),
		OutputPaths:      []string{"stdout"},
		ErrorOutputPaths: []string{"stderr"},
		EncoderConfig: zapcore.EncoderConfig{
			MessageKey:    "msg",
			LevelKey:      "lvl",
			TimeKey:       "t",
			CallerKey:     "call",
			StacktraceKey: "trace",
			EncodeLevel:   zapcore.CapitalColorLevelEncoder,
			EncodeTime:    zapcore.ISO8601TimeEncoder,
		},
		InitialFields: map[string]interface{}{
			"pid": os.Getpid(),
		},
	}

	// We might want to write to a file in the future
	productionCfg := zap.Config{
		Encoding:         "console",
		Level:            zap.NewAtomicLevelAt(level),
		OutputPaths:      []string{"stdout"},
		ErrorOutputPaths: []string{"stderr"},
		EncoderConfig: zapcore.EncoderConfig{
			MessageKey:    "message",
			LevelKey:      "level",
			TimeKey:       "timestamp",
			CallerKey:     "caller",
			StacktraceKey: "trace",
			EncodeLevel:   zapcore.LowercaseLevelEncoder,
			EncodeTime:    zapcore.ISO8601TimeEncoder,
		},
		InitialFields: map[string]interface{}{
			"pid": os.Getpid(),
		},
	}

	logger := zap.Must(productionCfg.Build())
	if os.Getenv("APP_ENV") == "development" {
		logger = zap.Must(developmentCfg.Build())
	}
	return logger
}

func init() {
	logger := createLogger()
	zap.ReplaceGlobals(logger)
	if logger.Core().Enabled(zap.DebugLevel) {
		zap.L().Debug("zap logger initialized")
	}
}
