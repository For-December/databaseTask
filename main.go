package main

import (
	"databaseTask/storage/database"
	"databaseTask/utils/excel"
	"databaseTask/utils/logger"
)

func main() {
	println(database.Client.Error)
	company, err := excel.ReadCompany()
	if err != nil {
		logger.Error(err)
		return
	}
	logger.Info(len(company))
	return
}
