package api

import "go-admin/app/user/router"

// init registers the user package router by appending router.InitRouter to AppRouters.
func init() {
	//注册路由
	AppRouters = append(AppRouters, router.InitRouter)
}