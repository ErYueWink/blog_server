package api

import "gvb_server/api/settings_api"

type ApiGroup struct {
	SettingsApi settings_api.SettingsApi // 系统配置
}

var ApiGroupApp = new(ApiGroup)
