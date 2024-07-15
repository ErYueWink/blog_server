package core

import (
	"fmt"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
	"gvb_server/global"
	"log"
	"time"
)

// Init_gorm 初始化数据库连接
func Init_gorm() *gorm.DB {
	if global.Config.Mysql.Host == "" {
		global.Log.Warning("mysql_host is not exist")
		return nil
	}
	dsn := global.Config.Mysql.Dsn()
	var mysqlLogger logger.Interface
	if global.Config.System.Env == "dev" {
		// dev环境显示所有日志
		mysqlLogger = logger.Default.LogMode(logger.Info)
	} else if global.Config.System.Env == "debug" {
		// debug模式只显示警告日志
		mysqlLogger = logger.Default.LogMode(logger.Error)
	} else {
		// 其他环境仅仅显示警告日志
		mysqlLogger = logger.Default.LogMode(logger.Warn)
	}
	// 连接数据库
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{
		Logger: mysqlLogger,
	})
	if err != nil {
		log.Fatalf(fmt.Sprintf("open mysql err:%v", err))
		return nil
	}
	sqlDB, _ := db.DB()
	sqlDB.SetMaxIdleConns(global.Config.Mysql.MaxIdleConns) // 最大空闲连接数
	sqlDB.SetMaxOpenConns(global.Config.Mysql.MaxOpenConns) // 最大可容纳
	sqlDB.SetConnMaxLifetime(time.Hour * 4)                 // 连接最大复用时间,不能超过mysql的wait_time
	global.Log.Info("mysql init success")
	return db
}
