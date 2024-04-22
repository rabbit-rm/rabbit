package zapToolkit

import (
	"time"

	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

func NewProduction(options ...zap.Option) (*zap.Logger, error) {
	cfg := zap.NewProductionConfig()
	// time encoder
	cfg.EncoderConfig.EncodeTime = timeEncoder
	return cfg.Build(options...)
}

func NewDevelopment(options ...zap.Option) (*zap.Logger, error) {
	cfg := zap.NewDevelopmentConfig()
	// time encoder
	cfg.EncoderConfig.EncodeTime = timeEncoder
	return cfg.Build(options...)
}

func timeEncoder(t time.Time, enc zapcore.PrimitiveArrayEncoder) {
	zapcore.TimeEncoderOfLayout("2006-01-02T15:04:05")(t, enc)
}
