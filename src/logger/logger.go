package logger

import (
	"go.uber.org/zap"
)

func InitLogger() (*zap.Logger, error) {
	lgr, err := zap.NewProduction()
	if err != nil {
		return nil, err
	}

	return lgr, nil
}
