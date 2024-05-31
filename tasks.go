package main

import (
	"databaseTask/constant/define"
	"databaseTask/dbmodels"
	"databaseTask/storage/database"
	"databaseTask/utils/logger"
)

func Task1() {
	logger.Debug("########## TASK1 ##########")
	type T struct {
		OrgCode           string `json:"orgCode"`
		OrgName           string `json:"orgName"`
		RegisteredAddress string `json:"registeredAddress"`
	}
	res := make([]T, 0)
	if err := database.Client.
		Raw(define.Task1).Scan(&res).Error; err != nil {
		logger.Error(err)
		return
	}
	logger.InfoF("一共 %d 条数据", len(res))

	logger.Debug("########## END ##########")

	return
}

func Task2() {

	res := make([]dbmodels.Executive, 0)
	if err := database.Client.
		Raw(define.Task2).Scan(&res).Error; err != nil {
		logger.Error(err)
		return
	}
	logger.Info(len(res))
	return
}
