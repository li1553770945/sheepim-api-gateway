package middleware

import (
	"context"
	"github.com/cloudwego/hertz/pkg/app"
	"github.com/cloudwego/hertz/pkg/common/hlog"
	"github.com/cloudwego/hertz/pkg/common/utils"
	"github.com/li1553770945/sheepim-api-gateway/biz/constant"
	userRpc "github.com/li1553770945/sheepim-user-service/kitex_gen/user"
)

func AdminMiddleware() app.HandlerFunc {
	return func(ctx context.Context, c *app.RequestContext) {
		// 从上下文中获取用户信息
		userInfo, ok := ctx.Value(constant.UserContextKey).(*userRpc.UserInfoResp)
		if !ok || userInfo == nil {
			hlog.CtxInfof(ctx, "用户未登录")
			c.JSON(200, utils.H{"code": constant.Unauthorized, "message": "用户未登录"})
			c.Abort()
			return
		}

		if *userInfo.Role != constant.AdminRole { // 假设 UserInfoResp 有 Role 字段
			hlog.CtxInfof(ctx, "非 admin 无权操作")
			c.JSON(200, utils.H{"code": constant.Unauthorized, "message": "无权进行此操作"})
			c.Abort()
			return
		}

		// 用户是管理员，继续处理请求
		c.Next(ctx)
	}
}
