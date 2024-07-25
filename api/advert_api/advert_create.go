package advert_api

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"gvb_server/global"
	"gvb_server/models"
	"gvb_server/utils/res"
)

type AdvertRequest struct {
	Title  string `json:"title" required:"true" msg:"请输入标题"`
	Href   string `json:"href" required:"true" msg:"跳转链接"`
	Images string `json:"images" required:"true" msg:"请输入图片链接"`
	IsShow bool   `json:"is_show" required:"true" msg:"图片是否显示"`
}

// AdvertCreateView 发布广告
func (AdvertApi) AdvertCreateView(c *gin.Context) {
	var cr AdvertRequest
	err := c.ShouldBindJSON(&cr)
	if err != nil {
		res.FailErrorCode(res.ArgumentError, c)
		return
	}
	var advertModel models.AdvertModel
	count := global.DB.Take(&advertModel, "title = ?", cr.Title).RowsAffected
	if count == 0 {
		res.FailWithMsg(fmt.Sprintf("标题为：%s的广告重复 请重新输入", cr.Title), c)
		return
	}
	// 广告入库
	err = global.DB.Create(&models.AdvertModel{
		Title:  cr.Title,
		Href:   cr.Href,
		Images: cr.Images,
		IsShow: cr.IsShow,
	}).Error
	if err != nil {
		res.FailWithMsg("广告入库失败", c)
		return
	}
	res.OkWithMsg("广告入库成功", c)
}
