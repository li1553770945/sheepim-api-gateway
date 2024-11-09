package trace

import (
	serverconfig "github.com/cloudwego/hertz/pkg/common/config"
	"github.com/hertz-contrib/obs-opentelemetry/provider"
	hertztracing "github.com/hertz-contrib/obs-opentelemetry/tracing"
	"github.com/li1553770945/sheepim-api-gateway/biz/infra/config"
)

func InitTrace(config *config.Config) (provider.OtelProvider, serverconfig.Option, *hertztracing.Config) {
	p := provider.NewOpenTelemetryProvider(
		provider.WithServiceName(config.ServerConfig.ServiceName),
		provider.WithExportEndpoint(config.OpenTelemetryConfig.Endpoint),
		provider.WithInsecure(),
	)
	print(config.ServerConfig.ServiceName)
	tracer, cfg := hertztracing.NewServerTracer()
	return p, tracer, cfg

}
