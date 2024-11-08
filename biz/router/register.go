// Code generated by hertz generator. DO NOT EDIT.

package router

import (
	"github.com/cloudwego/hertz/pkg/app/server"
	auth "github.com/li1553770945/sheepim-api-gateway/biz/router/auth"
	project "github.com/li1553770945/sheepim-api-gateway/biz/router/project"
	user "github.com/li1553770945/sheepim-api-gateway/biz/router/user"
)

// GeneratedRegister registers routers generated by IDL.
func GeneratedRegister(r *server.Hertz) {
	//INSERT_POINT: DO NOT DELETE THIS LINE!
	project.Register(r)

	auth.Register(r)

	user.Register(r)

}
