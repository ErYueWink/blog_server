package routers

import "gvb_server/api"

func (r RouterGroup) MenuRouter() {
	app := api.ApiGroupApp.MenuApi
	r.POST("/menu", app.MenuCreateView)
}
