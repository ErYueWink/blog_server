package routers

import "gvb_server/api"

func (r RouterGroup) MenuRouter() {
	app := api.ApiGroupApp.MenuApi
	r.POST("/menu", app.MenuCreateView)
	r.GET("/menu/:id", app.MenuDetailView)
	r.GET("/menu", app.MenuListView)
	r.DELETE("/menu", app.MenuRemoveView)
}
