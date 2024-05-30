package utils

import (
	"time"

	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

var Logger *zap.Logger

func InitLogger(mode string) {
	encoderCfg := zapcore.EncoderConfig{
		TimeKey:     "time",
		EncodeTime:  customTimeEncoder,
		MessageKey:  "msg",
		LevelKey:    "level",
		EncodeLevel: zapcore.LowercaseLevelEncoder,
	}

	var config zap.Config

	if mode == "dev" {
		// For development mode
		config = zap.NewDevelopmentConfig()
	} else {
		// For production mode
		config = zap.NewProductionConfig()
	}

	config.Level = zap.NewAtomicLevelAt(zap.InfoLevel)
	config.OutputPaths = []string{"stdout"}
	config.ErrorOutputPaths = []string{"stderr"}
	config.EncoderConfig = encoderCfg

	var err error
	Logger, err = config.Build()
	if err != nil {
		// Log the error or return it based on your application requirements
		// Example: log.Fatalf("create logger error: %v", err)
		panic("create logger error: " + err.Error())
	}
}

func customTimeEncoder(t time.Time, enc zapcore.PrimitiveArrayEncoder) {
	enc.AppendString(t.Format("2006-01-02 15:04:05"))
}
