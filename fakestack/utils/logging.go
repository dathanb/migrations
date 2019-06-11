package utils

import (
	"context"
	"github.com/sirupsen/logrus"
)

func Logger(ctx context.Context) *logrus.Entry {
	if logger, ok := ctx.Value("logger").(*logrus.Entry); ok {
		return logger
	}
	return logrus.WithField("request_id", "n/a")
}
