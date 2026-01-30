package router

import (
	"os"

	"github.com/gin-gonic/gin"
	log "github.com/go-admin-team/go-admin-core/logger"
	"github.com/go-admin-team/go-admin-core/sdk"
	common "go-admin/common/middleware"
)

// InitRouter 初始化路由并注册应用业务路由。
// 从 sdk.Runtime 获取 Gin 引擎，初始化认证中间件并调用 initRouter 注册路由；若未找到引擎、引擎类型不支持或 JWT 初始化失败，则记录致命日志并退出。
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