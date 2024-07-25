package advert_api

import (
	"github.com/gin-gonic/gin"
	"gvb_server/models"
	"gvb_server/service/common"
	"gvb_server/utils/res"
	"strings"
)

// AdvertListView 查询广告列表
func (AdvertApi) AdvertListView(c *gin.Context) {
	var cr models.PageInfo
	err := c.ShouldBindQuery(&cr)
	if err != nil {
		res.FailErrorCode(res.ArgumentError, c)
		return
	}
	// 网站用户默认只查询启用的广告，后台管理员可以查询所有广告
	isShow := true
	// 判断请求头中是否有admin 如果有：则显示所有广告
	refresh := c.GetHeader("GVB_REFERER")
	if strings.Contains(refresh, "admin") {
		isShow = false
	}
	// gorm框架中 布尔类型的值当作查询条件时，false会查询所有数据
	list, count, err := common.CommonList(models.AdvertModel{IsShow: isShow}, common.Option{
		PageInfo: cr,
		Debug:    true,
	})
	if err != nil {
		res.FailWithMsg("列表数据查询失败", c)
		return
	}
	res.OkWithList(list, count, c)
}
