package router

import (
	"github.com/gin-gonic/gin"
	jwt "github.com/go-admin-team/go-admin-core/sdk/pkg/jwtauth"

	"go-admin/app/user/apis"
	"go-admin/common/middleware"
)

// init 将 registerUserRouter 添加到 routerCheckRole 列表，以便在路由初始化阶段注册用户相关路由。
func init() {
	routerCheckRole = append(routerCheckRole, registerUserRouter)
}

// v1 是父级 gin.RouterGroup；authMiddleware 是用于路由组的 JWT 身份验证中间件。
func registerUserRouter(v1 *gin.RouterGroup, authMiddleware *jwt.GinJWTMiddleware) {
	api := apis.User{}
	r := v1.Group("/user").Use(authMiddleware.MiddlewareFunc()).Use(middleware.AuthCheckRole())
	{
		r.GET("", api.GetPage)
		r.GET("/:id", api.Get)
		r.POST("", api.Insert)
		r.PUT("/:id", api.Update)
		r.DELETE("", api.Delete)
	}
}