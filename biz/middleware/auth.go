package middleware

import (
	"context"
	"errors"
	"fmt"
	"github.com/cloudwego/hertz/pkg/app"
	"github.com/cloudwego/hertz/pkg/common/hlog"
	"github.com/cloudwego/hertz/pkg/common/utils"
	"github.com/li1553770945/sheepim-api-gateway/biz/constant"
	authRpc "github.com/li1553770945/sheepim-auth-service/kitex_gen/auth"
	"github.com/li1553770945/sheepim-auth-service/kitex_gen/auth/authservice"
	userRpc "github.com/li1553770945/sheepim-user-service/kitex_gen/user"
	"github.com/li1553770945/sheepim-user-service/kitex_gen/user/userservice"
	"strings"
	"sync"
)

// 定义上下文中的用户信息 key

// 将用户信息设置到上下文中的帮助函数
func withUserInfo(ctx context.Context, userInfo *userRpc.UserInfoResp) context.Context {
	return context.WithValue(ctx, constant.UserContextKey, userInfo)
}

// 从上下文中提取用户信息的帮助函数
func GetUserInfoFromCtx(ctx context.Context) (*userRpc.UserInfoResp, error) {
	userInfo, ok := ctx.Value(constant.UserContextKey).(*userRpc.UserInfoResp)
	if !ok {
		return nil, errors.New("未找到用户信息")
	}
	return userInfo, nil
}

var GlobalAuthMiddleware app.HandlerFunc
var once sync.Once

func GetGlobalGlobalAuthMiddleware() app.HandlerFunc {
	if GlobalAuthMiddleware == nil {
		panic("中间价在使用前未初始化")
	}
	return GlobalAuthMiddleware
}

func InitGlobalAuthMiddleware(authClient authservice.Client, userClient userservice.Client) {
	once.Do(func() {
		GlobalAuthMiddleware = AuthMiddleware(authClient, userClient)
	})
}

// AuthMiddleware 认证中间件
func AuthMiddleware(authClient authservice.Client, userClient userservice.Client) app.HandlerFunc {
	return func(ctx context.Context, c *app.RequestContext) {
		// 获取 Authorization 头部的 token
		token := string(c.GetHeader("Authorization"))
		if token == "" {
			c.JSON(200, utils.H{"code": constant.Unauthorized, "message": "您还未登陆，请先登录"})
			c.Abort()
			return
		}
		const bearerPrefix = "Bearer "
		if len(token) > len(bearerPrefix) && strings.HasPrefix(token, bearerPrefix) {
			token = token[len(bearerPrefix):] // 去除前缀，获取实际令牌
		}
		// 调用 Auth 服务验证 token 并获取用户 ID
		req := authRpc.GetUserIdReq{Token: token}
		resp, err := authClient.GetUserId(ctx, &req)
		if err != nil {
			hlog.CtxErrorf(ctx, "认证服务调用失败:%v", err)
			c.JSON(200, utils.H{"code": constant.SystemError, "message": "认证服务调用失败:" + err.Error()})
			c.Abort()
			return
		}
		if resp.BaseResp.Code != 0 {
			hlog.CtxInfof(ctx, "获取用户ID失败: %v，token 可能已失效", resp.BaseResp.Code)
			c.JSON(200, utils.H{"code": constant.Unauthorized, "message": fmt.Sprintf("获取用户ID失败: %v，token 可能已失效", resp.BaseResp.Message)})
			c.Abort()
			return
		}

		userId := resp.UserId
		userReq := userRpc.UserInfoReq{UserId: userId}
		userInfo, err := userClient.GetUserInfo(ctx, &userReq)

		if err != nil {
			hlog.CtxErrorf(ctx, "用户服务调用失败:%v", err)
			c.JSON(200, utils.H{"code": constant.SystemError, "message": "用户服务调用失败:" + err.Error()})
			c.Abort()
			return
		}

		if userInfo.BaseResp.Code != constant.Success {
			hlog.CtxInfof(ctx, "获取用户信息失败。状态码:%v", userInfo.BaseResp.Code)
			c.JSON(200, utils.H{"code": constant.SystemError, "message": fmt.Sprintf("获取用户信息失败。状态码:%v", userInfo.BaseResp.Code)})
			c.Abort()
			return
		}

		// 将用户信息添加到上下文中，供后续的处理器使用
		ctx = withUserInfo(ctx, userInfo)

		// 调用下一个中间件或处理器
		c.Next(ctx)
	}
}
