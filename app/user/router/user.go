package router

import (
	"github.com/gin-gonic/gin"
	jwt "github.com/go-admin-team/go-admin-core/sdk/pkg/jwtauth"

	"go-admin/app/user/apis"
	"go-admin/common/middleware"
)

// init appends registerUserRouter to routerCheckRole so the user routes are registered during router setup.
func init() {
	routerCheckRole = append(routerCheckRole, registerUserRouter)
}

// registerUserRouter registers user-related HTTP routes under the provided router group.
// The routes are mounted at "/user": GET "" for a paginated list, GET "/:id" to fetch a user by ID,
// POST "" to create a user, PUT "/:id" to update a user, and DELETE "" to remove a user.
// All routes are protected by the provided JWT authentication middleware and the role-check middleware.
// v1 is the parent Gin router group; authMiddleware supplies the JWT middleware used for authentication.
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