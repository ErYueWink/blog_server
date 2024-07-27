package menu_api

import (
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
	"gvb_server/global"
	"gvb_server/models"
	"gvb_server/utils/res"
)

// MenuRemoveView 删除菜单
func (MenuApi) MenuRemoveView(c *gin.Context) {
	var cr models.RemoveRequest
	err := c.ShouldBindJSON(&cr)
	if err != nil {
		res.FailErrorCode(res.ArgumentError, c)
		return
	}
	var menuList []models.MenuModel
	count := global.DB.Find(&menuList, cr.IDList).RowsAffected
	if count == 0 {
		res.FailWithMsg("菜单不存在", c)
		return
	}
	// 开启事务
	err = global.DB.Transaction(func(tx *gorm.DB) error {
		// 清空关联数据
		err := global.DB.Model(&menuList).Association("Banners").Clear()
		if err != nil {
			global.Log.Error(err.Error())
			return err
		}
		// 删除菜单数据
		err = global.DB.Delete(&menuList).Error
		if err != nil {
			global.Log.Error(err.Error())
			return err
		}
		return nil
	})
	if err != nil {
		res.FailWithMsg("删除菜单数据失败", c)
		return
	}
	res.OkWithMsg("删除菜单数据成功", c)
}
