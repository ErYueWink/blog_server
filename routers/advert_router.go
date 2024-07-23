package routers

import "gvb_server/api"

func (router RouterGroup) AdvertRouter() {
	app := api.ApiGroup{}.AdvertApi
	router.POST("/advert", app.AdvertCreateView)
}