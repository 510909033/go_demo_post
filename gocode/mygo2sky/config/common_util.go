package config

import (
	"context"
	"github.com/SkyAPM/go2sky"
	"time"
)

func GetLocalspan(ctx context.Context, tracer *go2sky.Tracer) (context.Context, func(err error)) {

	span, newCtx, err := tracer.CreateLocalSpan(ctx)
	if err != nil {
		return ctx, func(err error) {}
	}

	return newCtx, func(err error) {
		if err != nil {
			span.Error(time.Now(), err.Error())
		}
		span.End()
	}

}
