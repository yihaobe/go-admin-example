package api

import "go-admin/app/user/router"

// init 在包初始化阶段注册用户路由，将 router.InitRouter 添加到全局 AppRouters 列表以便在应用启动时初始化该路由。
func init() {
	//注册路由
	AppRouters = append(AppRouters, router.InitRouter)
}