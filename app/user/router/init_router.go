package router

import (
	"os"

	"github.com/gin-gonic/gin"
	log "github.com/go-admin-team/go-admin-core/logger"
	"github.com/go-admin-team/go-admin-core/sdk"
	common "go-admin/common/middleware"
)

// InitRouter initializes the application's Gin router by obtaining the engine from sdk.Runtime, ensuring it is a *gin.Engine, initializing authentication middleware, and registering business routes via initRouter.
// It logs a fatal error and exits the process if the engine is missing or of an unsupported type, or if authentication middleware initialization fails.
func InitRouter() {
	var r *gin.Engine
	h := sdk.Runtime.GetEngine()
	if h == nil {
		log.Fatal("not found engine...")
		os.Exit(-1)
	}
	switch h.(type) {
	case *gin.Engine:
		r = h.(*gin.Engine)
	default:
		log.Fatal("not support other engine")
		os.Exit(-1)
	}

	authMiddleware, err := common.AuthInit()
	if err != nil {
		log.Fatalf("JWT Init Error, %s", err.Error())
	}

	// 注册业务路由
	initRouter(r, authMiddleware)
}