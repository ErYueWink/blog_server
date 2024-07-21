package images_api

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"gvb_server/global"
	"gvb_server/models"
	"gvb_server/utils/res"
)

// ImagesRemoveView 批量删除图片
func (ImagesApi) ImagesRemoveView(c *gin.Context) {
	var cr models.RemoveRequest
	err := c.ShouldBindJSON(&cr)
	if err != nil {
		global.Log.Error(err.Error())
		res.FailWithMsg("参数绑定失败", c)
		return
	}
	var bannerList []models.BannerModel
	count := global.DB.Select("id").Find(&bannerList, cr).RowsAffected
	if count == 0 {
		global.Log.Error(err.Error())
		res.FailWithMsg(fmt.Sprintf("图片编号为%v的图片不存在", cr), c)
		return
	}
	err = global.DB.Delete(&bannerList).Error
	if err != nil {
		global.Log.Error(err.Error())
		res.FailWithMsg("删除图片失败", c)
		return
	}
	res.OkWithMsg(fmt.Sprintf("共删除%d张图片", cr), c)
	return
}
