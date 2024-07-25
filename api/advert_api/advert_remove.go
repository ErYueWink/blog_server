package advert_api

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"gvb_server/global"
	"gvb_server/models"
	"gvb_server/utils/res"
)

// AdvertRemoveView 删除广告
func (AdvertApi) AdvertRemoveView(c *gin.Context) {
	var cr models.RemoveRequest
	err := c.ShouldBindJSON(&cr)
	if err != nil {
		res.FailErrorCode(res.ArgumentError, c)
		return
	}
	var advertModels []models.AdvertModel
	count := global.DB.Find(&advertModels, "id = ?", cr.IDList).RowsAffected
	if count <= 0 { // 删除广告失败
		res.FailWithMsg(fmt.Sprintf("编号为%v的图片不存在", cr.IDList), c)
		return
	}
	// 删除广告
	err = global.DB.Delete(&advertModels).Error
	if err != nil {
		res.FailWithMsg("删除广告失败", c)
		return
	}
	res.OkWithMsg(fmt.Sprintf("共删除%d条广告", count), c)
}
