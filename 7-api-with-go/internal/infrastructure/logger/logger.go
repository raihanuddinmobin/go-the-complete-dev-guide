package logger

import (
	"context"

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

func TraceIdFromContext(c context.Context) string {
	traceID := c.Value("traceID")
	if traceID == "" {
		return ""
	}
	return traceID.(string)
}

func LoggerWithContext(ctx context.Context) *zap.Logger {
	traceID, _ := ctx.Value("traceID").(string)

	if traceID != "" {
		return Log.With(zap.String("traceID", traceID))
	}

	return Log
}
