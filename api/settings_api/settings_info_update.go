package settings_api

import (
	"errors"
	"fmt"
	"github.com/gin-gonic/gin"
	"gvb_server/config"
	"gvb_server/core"
	"gvb_server/global"
	"gvb_server/utils/res"
)

func (SettingsApi) SettingInfoUpdateView(c *gin.Context) {
	var cr SettingsUri
	err := c.ShouldBindUri(&cr)
	if err != nil {
		global.Log.Error("configuration file is not exist")
		res.FailWithMsg("configuration file is not exist", c)
		return
	}
	err = UpdateSettingsFile(cr.Name, c)
	if err != nil {
		res.FailWithMsg(fmt.Sprintf("%s the configuration file fails to be a modifed or does not exist"), c)
		return
	}
	res.OkWithMsg(fmt.Sprintf("%s the configuration file is successfully updated.", cr.Name), c)
}

// UpdateSettingsFile Modify different profiles based on the profile name
func UpdateSettingsFile(name string, c *gin.Context) error {
	switch name {
	case QQNAME:
		var qq config.QQ
		err := c.ShouldBindJSON(&qq)
		if err != nil {
			return err
		}
		global.Config.QQ = qq
		break
	case JWTNAME:
		var jwt config.Jwt
		err := c.ShouldBindJSON(&jwt)
		if err != nil {
			return err
		}
		global.Config.Jwt = jwt
		break
	case EMAILNAME:
		var email config.Email
		err := c.ShouldBindJSON(&email)
		if err != nil {
			return err
		}
		global.Config.Email = email
		break
	case QINIUNAME:
		var qiniu config.QiNiu
		err := c.ShouldBindJSON(&qiniu)
		if err != nil {
			return err
		}
		global.Config.QiNiu = qiniu
		break
	default:
		return errors.New(fmt.Sprintf("%s configuration not exist", name))
	}
	err := core.Set_yaml()
	if err != nil {
		return err
	}
	return nil
}
