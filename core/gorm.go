package core

import (
	"fmt"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
	"gwa_server/global"
	"gwa_server/models/system"
	"os"
	"time"
)

func InitGorm() *gorm.DB {
	if global.GWA_Config.Mysql.Host == "" {
		global.GWA_Log.Warnln("未配置mysql，取消gorm连接")
		return nil
	}
	dsn := global.GWA_Config.Mysql.Dsn()

	var mysqlLogger logger.Interface
	if global.GWA_Config.System.Env == "debug" {
		// 开发环境显示所有的sql
		mysqlLogger = logger.Default.LogMode(logger.Info)
	} else {
		mysqlLogger = logger.Default.LogMode(logger.Error) // 只打印错误的sql
	}
	//global.MysqlLog = logger.Default.LogMode(logger.Info)

	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{
		Logger: mysqlLogger,
	})
	if err != nil {
		global.GWA_Log.Fatalf(fmt.Sprintf("[%s] mysql连接失败", dsn))
	}
	sqlDB, _ := db.DB()
	sqlDB.SetMaxIdleConns(10)               // 最大空闲连接数
	sqlDB.SetMaxOpenConns(100)              // 最多可容纳
	sqlDB.SetConnMaxLifetime(time.Hour * 4) // 连接最大复用时间，不能超过mysql的wait_timeout

	return db
}

// 初始化数据库结构
func RegisterTables() {
	db := global.GWA_DB
	err := db.AutoMigrate(

		system.SysAdmin{},
		system.SysBaseMenu{},
		system.SysAuthority{},
		system.SysApi{},
	)
	if err != nil {
		global.GWA_Log.Error("register table failed", err)
		os.Exit(0)
	}
	global.GWA_Log.Info("register table success")
}
