package advert_api

import (
	"fmt"
	"github.com/fatih/structs"
	"github.com/gin-gonic/gin"
	"gvb_server/global"
	"gvb_server/models"
	"gvb_server/utils/res"
)

// AdvertUpdateView 修改广告
func (AdvertApi) AdvertUpdateView(c *gin.Context) {
	id := c.Param("id")
	var cr AdvertRequest
	err := c.ShouldBindJSON(&cr)
	if err != nil {
		res.FailErrorCode(res.ArgumentError, c)
		return
	}
	var advertModel models.AdvertModel
	count := global.DB.Take(&advertModel, "id = ?", id).RowsAffected
	if count == 0 {
		res.FailWithMsg(fmt.Sprintf("编号为%d的广告不存在", id), c)
		return
	}
	// 将结构体转为map
	advertMap := structs.Map(&cr)
	err = global.DB.Model(&advertModel).Updates(advertMap).Error
	if err != nil {
		res.FailWithMsg("广告修改失败", c)
		return
	}
	res.OkWithMsg(fmt.Sprintf("编号为%d的广告修改成功", id), c)
}
