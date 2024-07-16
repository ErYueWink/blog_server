package settings_api

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"gvb_server/global"
	"gvb_server/utils/res"
)

var (
	QQNAME    = "qq"
	JWTNAME   = "jwt"
	EMAILNAME = "email"
	QINIUNAME = "qiniu"
)

type SettingsUri struct {
	Name string `uri:"name" binding:"required"`
}

// SettingsInfoView 获取某一项的配置文件信息
// @Tags 系统管理
// @Summary 显示某一项的配置信息
// @Description 显示某一项的配置信息  site email qq qiniu jwt chat_group
// @Param name path string  true  "name"
// @Param token header string  true  "token"
// @Router /api/settings/{name} [get]
// @Produce json
// @Success 200 {object} res.Response{}
func (SettingsApi) SettingsInfoView(c *gin.Context) {
	var cr SettingsUri
	err := c.ShouldBindUri(&cr)
	if err != nil {
		global.Log.Error("parameter error or is not exists`")
		res.FailWithMsg("parameter error or is not exists", c)
		return
	}
	GetSettingInfo(cr.Name, c)
}

// GetSettingInfo Condense multiple configurations info a single configuration interface
func GetSettingInfo(name string, c *gin.Context) {
	switch name {
	case EMAILNAME:
		res.OkWithData(global.Config.Email, c)
		break
	case QQNAME:
		res.OkWithData(global.Config.QQ, c)
		break
	case JWTNAME:
		res.OkWithData(global.Config.Jwt, c)
		break
	case QINIUNAME:
		res.OkWithData(global.Config.QiNiu, c)
		break
	default:
		res.FailWithMsg(fmt.Sprintf("没有%s的配置文件", name), c)
	}
}
