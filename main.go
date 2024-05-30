package main

import (
	"databaseTask/storage/database"
	"databaseTask/utils/excel"
	"databaseTask/utils/logger"
)

func main() {
	Test()
	return
}

func Test() {
	println(database.Client.Error)
	company, err := excel.ReadExecutive()
	if err != nil {
		logger.Error(err)
		return
	}
	logger.Info(len(company))
}
