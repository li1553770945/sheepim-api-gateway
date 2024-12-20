// Code generated by hertz generator.

package main

import (
	"context"
	"github.com/cloudwego/hertz/pkg/app/server"
	"github.com/cloudwego/hertz/pkg/common/hlog"
	"github.com/hertz-contrib/obs-opentelemetry/provider"
	hertztracing "github.com/hertz-contrib/obs-opentelemetry/tracing"
	"github.com/li1553770945/sheepim-api-gateway/biz/infra/container"
	"github.com/li1553770945/sheepim-api-gateway/biz/middleware"
	"github.com/li1553770945/sheepim-api-gateway/global_middleware"
	"net"
	"os"
)

func main() {
	env := os.Getenv("ENV")
	if env == "" {
		env = "development"
	}
	container.InitGlobalContainer(env)
	App := container.GetGlobalContainer()

	defer func(p provider.OtelProvider, ctx context.Context) {
		err := p.Shutdown(ctx)
		if err != nil {
			hlog.Fatalf("server stopped with error:%s", err)
		}
	}(App.TraceStruct.Provider, context.Background())

	//初始化服务

	middleware.InitGlobalAuthMiddleware(App.AuthRpcClient, App.UserRpcClient)

	//初始化server
	addr, err := net.ResolveTCPAddr("tcp", App.Config.ServerConfig.ListenAddress)
	if err != nil {
		panic("设置监听地址出错")
	}

	h := server.Default(
		App.TraceStruct.Option,
		server.WithHostPorts(addr.String()),
	)

	h.Use(hertztracing.ServerMiddleware(App.TraceStruct.Config))
	h.Use(global_middleware.TraceIdMiddleware())
	register(h)
	h.Spin()
}
