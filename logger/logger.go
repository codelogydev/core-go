package logger

import "go.uber.org/zap"

var Log *zap.Logger

func init() {
	Log = zap.NewNop()
}

func Init() {
	Log, _ = zap.NewProduction()
}
