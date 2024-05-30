package database

import "C"
import (
	"databaseTask/constant/config"
	"databaseTask/dbmodels"
	"databaseTask/utils/logger"
	"fmt"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	gormLogger "gorm.io/gorm/logger"
)

var Client *gorm.DB

func init() {
	var err error
	var cfg gorm.Config
	cfg = gorm.Config{
		PrepareStmt: true,
		Logger:      gormLogger.Default.LogMode(gormLogger.Info),
		//NamingStrategy: schema.NamingStrategy{
		//	TablePrefix: "test",
		//},
		ConnPool: nil,
	}
	dsn := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8mb4&parseTime=True&loc=Local",
		config.EnvCfg.MysqlUser, config.EnvCfg.MysqlPassword,
		config.EnvCfg.MysqlHost, config.EnvCfg.MysqlPort, config.EnvCfg.MysqlDatabase)
	if Client, err = gorm.Open(mysql.Open(dsn), &cfg); err != nil {
		panic(err)
	}

	TableAutoMigrate()
}

func TableAutoMigrate() {
	if !config.EnvCfg.AutoMigrate {
		logger.Info("未启用迁移数据库")
		return
	}
	if err := Client.AutoMigrate(&dbmodels.Company{}, &dbmodels.Executive{},
		&dbmodels.University{}, &dbmodels.AlumniAssociation{}); err != nil {
		panic(err)
		return
	}

}
