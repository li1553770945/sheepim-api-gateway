// Code generated by hertz generator.

package user

import (
	"github.com/cloudwego/hertz/pkg/app"
	"github.com/li1553770945/sheepim-api-gateway/biz/middleware"
)

func rootMw() []app.HandlerFunc {
	// your code...
	return nil
}

func _getuserinfoMw() []app.HandlerFunc {
	// your code...
	return []app.HandlerFunc{
		middleware.GetGlobalGlobalAuthMiddleware(), // 验证用户是否已登录
		middleware.AdminMiddleware(),               // 检查用户是否具有管理员权限
	}
}

func _apiMw() []app.HandlerFunc {
	// your code...
	return nil
}

func _usersMw() []app.HandlerFunc {
	// your code...
	return nil
}

func _getselfinfoMw() []app.HandlerFunc {
	// your code...

	return []app.HandlerFunc{
		middleware.GetGlobalGlobalAuthMiddleware(), // 验证用户是否已登录
	}
}
