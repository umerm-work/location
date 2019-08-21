package service

import (
	"context"
	"fmt"
	"github.com/fb-investments/aone-location-service/pkg/io"
	"github.com/go-kit/kit/log"
)

// Middleware describes a service middleware.
type Middleware func(LocationService) LocationService

type loggingMiddleware struct {
	logger log.Logger
	next   LocationService
}

// LoggingMiddleware takes a logger as a dependency
// and returns a LocationService Middleware.
func LoggingMiddleware(logger log.Logger) Middleware {
	return func(next LocationService) LocationService {
		return &loggingMiddleware{logger, next}
	}

}

func (l loggingMiddleware) ListCountries(ctx context.Context) (res io.Response) {
	defer func() {
		l.logger.Log("method", "ListCountries", "res", res)
	}()
	return l.next.ListCountries(ctx)
}
func (l loggingMiddleware) CreateCountries(ctx context.Context, u io.Country) (res io.Response) {
	defer func() {
		l.logger.Log("method", "CreateCountries", "request", fmt.Sprintf("%+v\n", u), "success", res.Success, "internal error", res.Error.InternalMessage, "user message", res.Error.UserMessage)
	}()
	return l.next.CreateCountries(ctx, u)
}
func (l loggingMiddleware) ListStates(ctx context.Context, u io.UID) (res io.Response) {
	defer func() {
		l.logger.Log("method", "ListStates", "request", fmt.Sprintf("%+v\n", u), "success", res.Success, "internal error", res.Error.InternalMessage, "user message", res.Error.UserMessage)
	}()
	return l.next.ListStates(ctx, u)
}
func (l loggingMiddleware) CreateStates(ctx context.Context, u io.State) (res io.Response) {
	defer func() {
		l.logger.Log("method", "CreateStates", "request", fmt.Sprintf("%+v\n", u), "success", res.Success, "internal error", res.Error.InternalMessage, "user message", res.Error.UserMessage)
	}()
	return l.next.CreateStates(ctx, u)
}
func (l loggingMiddleware) ListCity(ctx context.Context, u io.UID) (res io.Response) {
	defer func() {
		l.logger.Log("method", "ListCity", "request", fmt.Sprintf("%+v\n", u), "success", res.Success, "internal error", res.Error.InternalMessage, "user message", res.Error.UserMessage)
	}()
	return l.next.ListCity(ctx, u)
}
func (l loggingMiddleware) CreateCity(ctx context.Context, c io.City) (res io.Response) {
	defer func() {
		l.logger.Log("method", "CreateCity", "request", fmt.Sprintf("%+v\n", c), "success", res.Success, "internal error", res.Error.InternalMessage, "user message", res.Error.UserMessage)
	}()
	return l.next.CreateCity(ctx, c)
}
