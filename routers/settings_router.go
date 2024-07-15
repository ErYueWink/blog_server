package routers

import "gvb_server/api"

func (router RouterGroup) SettingsRouter() {
	api := api.ApiGroupApp.SettingsApi
	router.GET("/settings_info", api.SettingsInfoView)
}
