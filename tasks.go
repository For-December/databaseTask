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
	logger.Debug("########## TASK2 ##########")
	res := make([]dbmodels.Executive, 0)
	if err := database.Client.
		Raw(define.Task2).Scan(&res).Error; err != nil {
		logger.Error(err)
		return
	}
	logger.Info(len(res))
	logger.Debug("########## END ##########")
	return
}

// Task2Map 其他 Task 我就不一一写结构体了，结果直接加入到map中
func Task2Map(rawSql string, num int) {
	logger.DebugF("########## TASK%d ##########", num)

	res := make(map[string]interface{})
	if err := database.Client.
		Raw(rawSql).Scan(&res).Error; err != nil {
		logger.Error(err)
		return
	}
	logger.Info(len(res))
	logger.Debug("########## END ##########")

	return
}

func Task3() {
	Task2Map(define.Task3, 3)
}

func Task4() {
	Task2Map(define.Task4, 4)
}
func Task5() {
	Task2Map(define.Task5, 5)
}
func Task6() {
	Task2Map(define.Task6, 6)
}
func Task7() {
	Task2Map(define.Task7, 7)
}
func Task8() {
	Task2Map(define.Task8, 8)
}
func Task9() {
	Task2Map(define.Task9, 9)
}
func Task10() {
	Task2Map(define.Task10, 10)
}
func Task11() {
	Task2Map(define.Task11, 11)
}
func Task12() {
	Task2Map(define.Task12, 12)
}
func Task13() {
	Task2Map(define.Task13, 13)
}
func Task14() {
	Task2Map(define.Task14, 14)
}
func Task15() {
	Task2Map(define.Task15, 15)
}
func Task16() {
	Task2Map(define.Task16, 16)
}
func Task17() {
	logger.Warning("这个应该在数据库中执行")
	//Task2Map(define.Task17, 17)
}
