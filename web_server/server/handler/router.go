package handler

import (
	"vue-admin/web_server/server/handler/dev_tools_handler"
	"vue-admin/web_server/server/handler/handler_common"
	"vue-admin/web_server/server/handler/index_handler"
	"vue-admin/web_server/server/handler/log_handelr"
	"vue-admin/web_server/server/handler/menu_handler"
	"vue-admin/web_server/server/handler/monitor_handler"
	"vue-admin/web_server/server/handler/permission_handler"
	"vue-admin/web_server/server/handler/system_handler"
	"vue-admin/web_server/server/handler/verify_handler"

	"vue-admin/web_server/server/handler/app_handler"
	"vue-admin/web_server/server/handler/file_handler"
	"vue-admin/web_server/server/handler/role_handler"
	"vue-admin/web_server/server/handler/user_handler"

	"strings"
	"vue-admin/web_server/server/middleware"

	"github.com/gin-gonic/gin"
)

var (
	routerList []handler_common.GinHandleFunc
)

//host:port/auth_prefix/prefix/path
func RegisterRouters(r *gin.Engine, prefix string, authPrefix string) {
	jwtR := r.Group(prefix + authPrefix)
	jwtR.Use(middleware.JWTAuthMiddleware())
	addAllRouters()

	for _, v := range routerList {
		route := strings.ToLower(v.Route)
		switch v.RouterType {
		case handler_common.RouterTypeNeedAuth:
			funcDoRouteNeedAuthRegister(strings.ToUpper(v.Method), route, v.Handler, jwtR)
		case handler_common.RouterTypeNormal:
			route = strings.ToLower(prefix + route)
			funcDoRouteRegister(strings.ToUpper(v.Method), route, v.Handler, r)
		}
	}
}

//使用jwt过滤的routerType==routerTypeNeedAuth
func funcDoRouteNeedAuthRegister(method, route string, handler gin.HandlerFunc, jwtR *gin.RouterGroup) {
	switch method {
	case "POST":
		jwtR.POST(route, handler)
	case "PUT":
		jwtR.PUT(route, handler)
	case "HEAD":
		jwtR.HEAD(route, handler)
	case "DELETE":
		jwtR.DELETE(route, handler)
	case "OPTIONS":
		jwtR.OPTIONS(route, handler)
	default:
		jwtR.GET(route, handler)
	}
}

//普通routerType==routerTypeNormal
func funcDoRouteRegister(method, route string, handler gin.HandlerFunc, r *gin.Engine) {
	switch method {
	case "POST":
		r.POST(route, handler)
	case "PUT":
		r.PUT(route, handler)
	case "HEAD":
		r.HEAD(route, handler)
	case "DELETE":
		r.DELETE(route, handler)
	case "OPTIONS":
		r.OPTIONS(route, handler)
	default:
		r.GET(route, handler)
	}
}

//添加route 到RouterList
func addAllRouters() {
	routerList = append(routerList, user_handler.Routers...)
	routerList = append(routerList, file_handler.Routers...)
	routerList = append(routerList, app_handler.Routers...)
	routerList = append(routerList, role_handler.Routers...)
	routerList = append(routerList, permission_handler.Routers...)
	routerList = append(routerList, verify_handler.Routers...)
	routerList = append(routerList, menu_handler.Routers...)
	routerList = append(routerList, log_handelr.Routers...)
	routerList = append(routerList, system_handler.Routers...)
	routerList = append(routerList, index_handler.Routers...)
	routerList = append(routerList, monitor_handler.Routers...)
	routerList = append(routerList, dev_tools_handler.Routers...)
}
