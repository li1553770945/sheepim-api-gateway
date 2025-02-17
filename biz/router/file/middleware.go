// Code generated by hertz generator.

package file

import (
	"github.com/cloudwego/hertz/pkg/app"
	"github.com/li1553770945/sheepim-api-gateway/biz/middleware"
)

func rootMw() []app.HandlerFunc {
	// your code...
	return nil
}

func _apiMw() []app.HandlerFunc {
	// your code...
	return nil
}

func _uploadfileMw() []app.HandlerFunc {
	return []app.HandlerFunc{
		middleware.GetGlobalGlobalAuthMiddleware(), // 验证用户是否已登录
	}
}

func _downloadfilereqMw() []app.HandlerFunc {
	// your code...
	return nil
}

func _deletefileMw() []app.HandlerFunc {
	return []app.HandlerFunc{
		middleware.GetGlobalGlobalAuthMiddleware(), // 验证用户是否已登录
	}
}

func _downloadfileMw() []app.HandlerFunc {
	// your code...
	return nil
}
