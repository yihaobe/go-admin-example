package router

import (
	"github.com/gin-gonic/gin"
	jwt "github.com/go-admin-team/go-admin-core/sdk/pkg/jwtauth"
)

var (
	routerNoCheckRole = make([]func(*gin.RouterGroup), 0)
	routerCheckRole   = make([]func(v1 *gin.RouterGroup, authMiddleware *jwt.GinJWTMiddleware), 0)
)

// initRouter initializes application routes on the provided Gin engine by
// registering routes that do not require authentication and routes that do,
// using the supplied JWT middleware. It returns the same *gin.Engine instance.
func initRouter(r *gin.Engine, authMiddleware *jwt.GinJWTMiddleware) *gin.Engine {

	// 无需认证的路由
	noCheckRoleRouter(r)
	// 需要认证的路由
	checkRoleRouter(r, authMiddleware)

	return r
}

// It invokes each function in routerNoCheckRole with the created group on the provided Gin engine.
func noCheckRoleRouter(r *gin.Engine) {
	// 可根据业务需求来设置接口版本
	v1 := r.Group("/api/v1")

	for _, f := range routerNoCheckRole {
		f(v1)
	}
}

// checkRoleRouter creates the "/api/v1" route group and registers routes that require authentication using the provided JWT middleware.
// It invokes each registered route installer with the created group and the authMiddleware.
func checkRoleRouter(r *gin.Engine, authMiddleware *jwt.GinJWTMiddleware) {
	// 可根据业务需求来设置接口版本
	v1 := r.Group("/api/v1")

	for _, f := range routerCheckRole {
		f(v1, authMiddleware)
	}
}