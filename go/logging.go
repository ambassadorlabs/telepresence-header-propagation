package main

import (
	"context"
	"time"

	"github.com/go-kit/kit/log"
)

func loggingMiddleware(logger log.Logger) ServiceMiddleware {
	return func(next StringService) StringService {
		return logmw{logger, next}
	}
}

type logmw struct {
	logger log.Logger
	StringService
}

func (mw logmw) InitialUppercase(ctx context.Context, s string) (output string, err error) {
	defer func(begin time.Time) {
		_ = mw.logger.Log(
			"method", "uppercase",
			"telepresenceHeader", ctx.Value("x-telepresence-id").(string),
		)
	}(time.Now())

	output, err = mw.StringService.InitialUppercase(ctx, s)
	return
}

func (mw logmw) FinalUppercase(ctx context.Context, s string) (output string, err error) {
	defer func(begin time.Time) {
		_ = mw.logger.Log(
			"method", "finalUppercase",
			"telepresenceHeader", ctx.Value("x-telepresence-id").(string),
		)
	}(time.Now())

	output, err = mw.StringService.FinalUppercase(ctx, s)
	return
}

func (mw logmw) Count(ctx context.Context, s string) (n int) {
	defer func(begin time.Time) {
		_ = mw.logger.Log(
			"method", "count",
			"input", s,
			"n", n,
			"took", time.Since(begin),
		)
	}(time.Now())

	n = mw.StringService.Count(ctx, s)
	return
}
