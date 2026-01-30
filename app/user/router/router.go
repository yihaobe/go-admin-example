package router

import (
	"github.com/gin-gonic/gin"
	jwt "github.com/go-admin-team/go-admin-core/sdk/pkg/jwtauth"
)

var (
	routerNoCheckRole = make([]func(*gin.RouterGroup), 0)
	routerCheckRole   = make([]func(v1 *gin.RouterGroup, authMiddleware *jwt.GinJWTMiddleware), 0)
)

// initRouter 初始化并配置传入的 Gin 引擎，在 /api/v1 下注册无需认证与需要认证的路由，并返回配置后的引擎。
func initRouter(r *gin.Engine, authMiddleware *jwt.GinJWTMiddleware) *gin.Engine {

	// 无需认证的路由
	noCheckRoleRouter(r)
	// 需要认证的路由
	checkRoleRouter(r, authMiddleware)

	return r
}

// noCheckRoleRouter 注册无需认证的路由到 /api/v1 路由组。
// 它在给定的 Gin 引擎上创建 "/api/v1" 路由组，并将 routerNoCheckRole 切片中的每个注册函数应用到该组以完成路由注册。
func noCheckRoleRouter(r *gin.Engine) {
	// 可根据业务需求来设置接口版本
	v1 := r.Group("/api/v1")

	for _, f := range routerNoCheckRole {
		f(v1)
	}
}

// checkRoleRouter 在给定 Gin 引擎上注册需要认证的路由，使用提供的 JWT 中间件。
// 它在 /api/v1 路径下创建路由组，并将 routerCheckRole 中的注册函数逐一应用到该组，传入 authMiddleware。
func checkRoleRouter(r *gin.Engine, authMiddleware *jwt.GinJWTMiddleware) {
	// 可根据业务需求来设置接口版本
	v1 := r.Group("/api/v1")

	for _, f := range routerCheckRole {
		f(v1, authMiddleware)
	}
}