package logger

import (
	"go.uber.org/zap"
	"mobin.dev/pkg/config"
	"mobin.dev/pkg/constants"
)

var Log *zap.Logger

func Init() error {
	env := config.Get().AppEnv
	var err error

	if env == constants.DEVELOPMENT {
		Log, err = zap.NewDevelopment()
	} else {
		Log, err = zap.NewProduction()
	}

	return err
}
