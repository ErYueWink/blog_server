package routers

import "gvb_server/api"

func (router RouterGroup) SettingsRouter() {
	app := api.ApiGroupApp.SettingsApi
	router.GET("/settings/:name", app.SettingsInfoView)
	router.PUT("/settings/:name", app.SettingInfoUpdateView)
}
