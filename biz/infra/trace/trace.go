package trace

import (
	serverconfig "github.com/cloudwego/hertz/pkg/common/config"
	"github.com/hertz-contrib/obs-opentelemetry/provider"
	hertztracing "github.com/hertz-contrib/obs-opentelemetry/tracing"
	"github.com/li1553770945/sheepim-api-gateway/biz/infra/config"
)

type TraceSturct struct {
	Provider provider.OtelProvider
	Option   serverconfig.Option
	Config   *hertztracing.Config
}

func InitTrace(config *config.Config) *TraceSturct {
	p := provider.NewOpenTelemetryProvider(
		provider.WithServiceName(config.ServerConfig.ServiceName),
		provider.WithExportEndpoint(config.OpenTelemetryConfig.Endpoint),
		provider.WithInsecure(),
	)
	print(config.ServerConfig.ServiceName)
	tracer, cfg := hertztracing.NewServerTracer()
	return &TraceSturct{
		Provider: p,
		Option:   tracer,
		Config:   cfg,
	}

}
