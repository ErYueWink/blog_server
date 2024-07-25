package routers

import "gvb_server/api"

func (router RouterGroup) AdvertRouter() {
	app := api.ApiGroup{}.AdvertApi
	router.POST("/advert", app.AdvertCreateView)
	router.GET("/advert", app.AdvertListView)
	router.DELETE("/advert", app.AdvertRemoveView)
	router.PUT("/advert/:id", app.AdvertUpdateView)
}
