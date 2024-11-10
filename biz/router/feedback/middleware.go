// Code generated by hertz generator.

package feedback

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

func _feedbackMw() []app.HandlerFunc {
	// your code...
	return nil
}

func _getfeedbackMw() []app.HandlerFunc {
	// your code...
	return nil
}

func _getreplyMw() []app.HandlerFunc {
	// your code...
	return nil
}

func _addreplyMw() []app.HandlerFunc {
	return []app.HandlerFunc{
		middleware.GetGlobalGlobalAuthMiddleware(), // 验证用户是否已登录
		middleware.AdminMiddleware(),               // 检查用户是否具有管理员权限
	}
}

func _addfeedbackMw() []app.HandlerFunc {
	// your code...
	return nil
}

func _feedback0Mw() []app.HandlerFunc {
	// your code...
	return nil
}

func _getfeedbackcategoriesMw() []app.HandlerFunc {
	// your code...
	return nil
}

func _getunreadfeedbackMw() []app.HandlerFunc {
	return []app.HandlerFunc{
		middleware.GetGlobalGlobalAuthMiddleware(), // 验证用户是否已登录
		middleware.AdminMiddleware(),               // 检查用户是否具有管理员权限
	}
}
