package menu_api

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"gvb_server/global"
	"gvb_server/models"
	"gvb_server/utils/res"
)

type Banner struct {
	ID   uint   `json:"id"`
	Path string `json:"path"`
}

type MenuResponse struct {
	models.MenuModel
	Banners []Banner `json:"banner"`
}

// MenuDetailView 查询菜单详情
func (MenuApi) MenuDetailView(c *gin.Context) {
	id := c.Param("id")
	var menuModel models.MenuModel
	count := global.DB.Take(&menuModel, "id = ?", id).RowsAffected
	if count == 0 {
		res.FailWithMsg(fmt.Sprintf("编号为%s的菜单不存在", id), c)
		return
	}
	// 查询中间表数据
	var menuBannerList []models.MenuBannerModel
	err := global.DB.Preload("BannerModel").Find(&menuBannerList, "menu_id = ?", id).Error
	if err != nil {
		res.FailWithMsg("查询中间表数据失败", c)
		return
	}
	var banners = make([]Banner, 0)
	for _, model := range menuBannerList {
		if menuModel.ID != model.MenuID {
			continue
		}
		banners = append(banners, Banner{
			ID:   model.BannerID,
			Path: model.BannerModel.Path,
		})
	}
	menuResponse := MenuResponse{
		MenuModel: menuModel,
		Banners:   banners,
	}
	res.OkWithData(menuResponse, c)
}
