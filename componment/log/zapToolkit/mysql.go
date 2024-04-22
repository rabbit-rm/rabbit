package zapToolkit

import (
	"go.uber.org/zap"
	"gorm.io/gorm/logger"
)

type mysqlLogger struct {
	output *zap.SugaredLogger
}

func (w *mysqlLogger) Printf(format string, args ...interface{}) {
	w.output.Infof(format, args)
}

func NewZapMySQL(w *zap.Logger) logger.Writer {
	// zapToolkit
	return &mysqlLogger{output: w.Sugar()}
}
