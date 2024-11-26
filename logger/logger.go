package logger

import (
	"go.uber.org/zap"
)

var Logger *zap.Logger

func InitLogger() (*zap.Logger, error) {
	var err error
	lgr, err := zap.NewProduction()
	if err != nil {
		return nil, err
	}
	
	return lgr, nil
}
