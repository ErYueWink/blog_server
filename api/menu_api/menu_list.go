package menu_api

import (
	"github.com/gin-gonic/gin"
	"gvb_server/global"
	"gvb_server/models"
	"gvb_server/utils/res"
)

// MenuListView 查询菜单列表
func (MenuApi) MenuListView(c *gin.Context) {
	var menuList []models.MenuModel
	var menuIDList []uint
	err := global.DB.Order("sort desc").Find(&menuList).Select("id").Scan(&menuIDList).Error
	if err != nil {
		res.FailWithMsg("查询菜单列表数据失败", c)
		return
	}
	// 查询中间表数据
	var menuBannerList []models.MenuBannerModel
	err = global.DB.Preload("BannerModel").Order("sort desc").Find(&menuBannerList, "menu_id in ?", menuIDList).Error
	if err != nil {
		res.FailWithMsg("中间表数据查询失败", c)
		return
	}
	var menus = make([]MenuResponse, 0)
	for _, menu := range menuList {
		var banners = make([]Banner, 0)
		for _, bannerModel := range menuBannerList {
			if menu.ID != bannerModel.MenuID {
				continue
			}
			banners = append(banners, Banner{
				ID:   bannerModel.BannerID,
				Path: bannerModel.BannerModel.Path,
			})
		}
		menus = append(menus, MenuResponse{
			MenuModel: menu,
			Banners:   banners,
		})
	}
	res.OkWithList(menus, int64(len(menus)), c)
}
