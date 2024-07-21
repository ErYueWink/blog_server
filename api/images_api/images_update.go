package images_api

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"gvb_server/global"
	"gvb_server/models"
	"gvb_server/utils/res"
)

type ImagesUpdateRequest struct {
	ID   uint   `json:"id" required:"true" msg:"请输入图片ID"`
	Name string `json:"name" required:"true" msg:"请输入图片名称"`
}

// ImagesUpdateView 修改图片名称
func (ImagesApi) ImagesUpdateView(c *gin.Context) {
	var cr ImagesUpdateRequest
	err := c.ShouldBindJSON(&cr)
	if err != nil {
		res.FailWithMsg("参数绑定失败", c)
		return
	}
	var bannerModel models.BannerModel
	count := global.DB.Take(&bannerModel, "id = ?", cr.ID).RowsAffected
	if count == 0 {
		res.FailWithMsg(fmt.Sprintf("id为%d的图片不存在", cr.ID), c)
		return
	}
	// 修改图片名称
	err = global.DB.Model(&bannerModel).Update("name", cr.Name).Error
	if err != nil {
		res.FailWithMsg("修改图片名称失败", c)
		return
	}
	res.OkWithMsg(fmt.Sprintf("id为%d的图片修改成功", cr.ID), c)
}
