package routers

import "gvb_server/api"

func (router RouterGroup) ImagesRouter() {
	app := api.ApiGroupApp.ImagesApi
	router.POST("/images", app.ImagesUploadView)
	router.GET("/images", app.ImagesListView)
	router.DELETE("/images", app.ImagesRemoveView)
	router.PUT("/images", app.ImagesUpdateView)
}
