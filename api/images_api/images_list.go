package images_api

import (
	"github.com/gin-gonic/gin"
	"gvb_server/models"
	"gvb_server/service/common"
	"gvb_server/utils/res"
)

// ImagesListView 图片列表查询
func (ImagesApi) ImagesListView(c *gin.Context) {
	var cr models.PageInfo
	err := c.ShouldBindQuery(&cr)
	if err != nil {
		res.FailWithMsg("参数绑定失败", c)
		return
	}
	list, count, err := common.CommonList[models.BannerModel](models.BannerModel{}, common.Option{
		PageInfo: cr,
		Debug:    true,
	})
	if err != nil {
		res.FailWithMsg("图片列表查询失败", c)
		return
	}
	res.OkWithList[[]models.BannerModel](list, count, c)

}
