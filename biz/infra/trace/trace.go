package trace

import (
	serverconfig "github.com/cloudwego/hertz/pkg/common/config"
	"github.com/hertz-contrib/obs-opentelemetry/provider"
	hertztracing "github.com/hertz-contrib/obs-opentelemetry/tracing"
	"sheepim-api-gateway/biz/infra/config"
)

func InitTrace(config *config.Config) (provider.OtelProvider, serverconfig.Option, *hertztracing.Config) {
	p := provider.NewOpenTelemetryProvider(
		provider.WithServiceName(config.ServerConfig.ServiceName),
		provider.WithExportEndpoint(config.OpenTelemetryConfig.Endpoint),
		provider.WithInsecure(),
	)
	tracer, cfg := hertztracing.NewServerTracer()
	return p, tracer, cfg

}
