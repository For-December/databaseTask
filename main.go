package main

import (
	"databaseTask/storage/database"
	"databaseTask/utils/excel"
	"databaseTask/utils/logger"
)

func main() {
	// 从excel文件中读取数据，写入数据库
	//InitDatabase()
	Task1()
	//Test()
}

func InitDatabase() {
	companies, err := excel.ReadCompany()
	if err != nil {
		logger.Error(err)
		return
	}
	if err = database.Client.Create(companies).Error; err != nil {
		logger.Error(err)
		return
	}

	universities, err := excel.ReadUniversity()
	if err != nil {
		logger.Error(err)
		return
	}
	if err = database.Client.Create(universities).Error; err != nil {
		logger.Error(err)
		return
	}

	executives, err := excel.ReadExecutive()
	if err != nil {
		logger.Error(err)
		return
	}
	logger.Info(len(executives))

	// 关闭外键约束
	database.Client.Exec("SET foreign_key_checks = 0")

	// 分成3组插入
	if err = database.Client.Create(executives[0 : len(executives)/3]).Error; err != nil {
		logger.Error(err)
		return
	}

	if err = database.Client.Create(executives[len(executives)/3 : 2*len(executives)/3]).Error; err != nil {
		logger.Error(err)
		return
	}

	if err = database.Client.Create(executives[2*len(executives)/3:]).Error; err != nil {
		logger.Error(err)
		return
	}

	// 启用外键约束
	database.Client.Exec("SET foreign_key_checks = 1")
}

func Test() {
	println(database.Client.Error)

	companies, err := excel.ReadCompany()
	if err != nil {
		logger.Error(err)
		return
	}
	logger.InfoF("%v 共有%d条数据", "companies", len(companies))

	executives, err := excel.ReadExecutive()
	if err != nil {
		logger.Error(err)
		return
	}
	logger.InfoF("%v 共有%d条数据", "executives", len(executives))

	universities, err := excel.ReadUniversity()
	if err != nil {
		logger.Error(err)
		return
	}
	logger.InfoF("%v 共有%d条数据", "universities", len(universities))
}
