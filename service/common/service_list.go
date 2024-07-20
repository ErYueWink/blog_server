package common

import (
	"gorm.io/gorm"
	"gvb_server/global"
	"gvb_server/models"
)

type Option struct {
	models.PageInfo
	Debug bool
}

// CommonList Mysql通用查询方法
func CommonList[T any](model T, option Option) (list []T, count int64, err error) {
	DB := global.DB
	if option.Debug {
		DB = global.DB.Session(&gorm.Session{Logger: global.MysqlLog})
	}
	// 排序默认值
	if option.Sort == "" {
		option.Sort = "create_at desc" // 默认根据时间倒序排序
	}
	// 图片总条数
	count = DB.Select("id").Find(&list).RowsAffected
	// 计算偏移量
	offset := (option.Page - 1) * option.Limit
	if offset < 0 {
		offset = 0
	}
	err = DB.Limit(option.Limit).Offset(offset).Order(option.Sort).Find(&list).Error
	return list, count, err
}
