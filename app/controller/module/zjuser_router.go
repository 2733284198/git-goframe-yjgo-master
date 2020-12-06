// ==========================================================================
// 云捷GO自动生成路由代码，只生成一次，按需修改,再次生成不会覆盖.
// 生成日期：2020-12-07 01:38:53
// 生成路径: app/controller/module/zjuser_router.go
// 生成人：yunjie
// ==========================================================================
package module

import (
	"yj-app/app/controller/module/zjuser"
	"yj-app/app/service/middleware/auth"
	"yj-app/app/service/middleware/router"
)

//加载路由
func init() {
	g1 := router.New("admin", "/module/zjuser", auth.Auth)
	g1.GET("/", "module:zjuser:view", zjuser.List)
	g1.POST("/list", "module:zjuser:list", zjuser.ListAjax)
	g1.GET("/add", "module:zjuser:add", zjuser.Add)
	g1.POST("/add", "module:zjuser:add", zjuser.AddSave)
	g1.POST("/remove", "module:zjuser:remove", zjuser.Remove)
	g1.GET("/edit", "module:zjuser:edit", zjuser.Edit)
	g1.POST("/edit", "module:zjuser:edit", zjuser.EditSave)
	g1.POST("/export", "module:zjuser:export", zjuser.Export)
}
