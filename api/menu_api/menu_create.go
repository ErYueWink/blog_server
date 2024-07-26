package menu_api

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"gvb_server/global"
	"gvb_server/models"
	"gvb_server/models/ctype"
	"gvb_server/utils/res"
)

type ImageSort struct {
	ImageID uint `json:"image_id"`
	Sort    int  `json:"sort"`
}

type MenuRequest struct {
	Title         string      `json:"title" binding:"required" msg:"请完善菜单名称" structs:"title"`
	Path          string      `json:"path" binding:"required" msg:"请完善菜单路径" structs:"path"`
	Slogan        string      `json:"slogan" structs:"slogan"`
	Abstract      ctype.Array `json:"abstract" structs:"abstract"`
	AbstractTime  int         `json:"abstract_time" structs:"abstract_time"` // 切换的时间，单位秒
	BannerTime    int         `json:"banner_time" structs:"banner_time"`     // 切换的时间，单位秒
	Sort          int         `json:"sort" structs:"sort"`                   // 菜单的序号
	ImageSortList []ImageSort `json:"image_sort_list" structs:"-"`           // 具体图片的顺序
}

// MenuCreateView 发布菜单
func (MenuApi) MenuCreateView(c *gin.Context) {
	var cr MenuRequest
	err := c.ShouldBindJSON(&cr)
	if err != nil {
		res.FailErrorCode(res.ArgumentError, c)
		return
	}
	// 查询菜单是否重复
	var menuModels []models.MenuModel
	count := global.DB.Find(&menuModels, "title = ? or path = ?", cr.Title, cr.Path).RowsAffected
	if count > 0 {
		res.FailWithMsg(fmt.Sprintf("标题为%s的菜单重复", cr.Title), c)
		return
	}
	menuModel := models.MenuModel{
		Title:        cr.Title,
		Path:         cr.Path,
		Slogan:       cr.Slogan,
		Abstract:     cr.Abstract,
		AbstractTime: cr.AbstractTime,
		BannerTime:   cr.BannerTime,
		Sort:         cr.Sort,
	}
	// 菜单入库
	err = global.DB.Create(&menuModel).Error
	if err != nil {
		res.FailWithMsg("图片入库失败", c)
		return
	}
	// 判断菜单有没有上传图片
	if len(cr.ImageSortList) == 0 {
		res.OkWithMsg("菜单入库成功", c)
		return
	}
	var menuBannerModels []models.MenuBannerModel
	for _, menuBanner := range cr.ImageSortList {
		// 判断图片是否存在
		var bannerModel models.BannerModel
		err = global.DB.Take(&bannerModel, "id = ?", menuBanner.ImageID).Error
		if err != nil {
			global.Log.Error(err.Error())
			continue
		}
		menuBannerModels = append(menuBannerModels, models.MenuBannerModel{
			MenuID:   menuModel.ID,
			BannerID: menuBanner.ImageID,
			Sort:     menuBanner.Sort,
		})
	}
	err = global.DB.Create(&menuBannerModels).Error
	if err != nil {
		res.FailWithMsg("图片数据关联失败", c)
		return
	}
	res.OkWithMsg("菜单数据入库成功", c)
}
