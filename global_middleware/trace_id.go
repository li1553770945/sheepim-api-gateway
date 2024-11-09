package global_middleware

import (
	"context"
	"github.com/cloudwego/hertz/pkg/app"
	oteltrace "go.opentelemetry.io/otel/trace"
)

func TraceIdMiddleware() app.HandlerFunc {
	return func(ctx context.Context, c *app.RequestContext) {
		// pre-handle
		// ...
		s := oteltrace.SpanFromContext(ctx).SpanContext().TraceID().String()
		c.Header("X-Trace-Id", s)
		c.Next(ctx)
	}
}
