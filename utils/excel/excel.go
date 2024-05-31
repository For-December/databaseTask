package excel

import (
	"databaseTask/dbmodels"
	"databaseTask/utils/calc"
	"databaseTask/utils/logger"
	"errors"
	"fmt"
	"github.com/xuri/excelize/v2"
	"reflect"
	"strconv"
)

var usefulCompanyCols = []string{
	"ORG_CODE", "ORG_NAME", "REG_ADDRESS",
	"EMP_NUM", "REG_CAPITAL", "INDUSTRYCSRC1", "TRADE_MARKET",
}

var companyExcel2Struct = map[string]string{
	"ORG_CODE":      "OrgCode",
	"ORG_NAME":      "OrgName",
	"REG_ADDRESS":   "RegisteredAddress",
	"EMP_NUM":       "EmployeeCount",
	"REG_CAPITAL":   "RegisteredCapital",
	"INDUSTRYCSRC1": "Industry",
	"TRADE_MARKET":  "TradeMarket",
}

func ReadCompany() (resInfos []dbmodels.Company, err error) {
	return readExcel[dbmodels.Company](
		"resources/上市公司列表.xlsx",
		usefulCompanyCols, companyExcel2Struct)
}

var usefulUniversityCols = []string{
	"序号", "学校名称", "省份", "城市",
}

var universityExcel2Struct = map[string]string{
	"序号":   "UniversityID",
	"学校名称": "UniversityName",
	"省份":   "Province",
	"城市":   "City",
}

func ReadUniversity() (resInfos []dbmodels.University, err error) {
	return readExcel[dbmodels.University](
		"resources/全国985大学.xlsx",
		usefulUniversityCols, universityExcel2Struct)
}

var usefulExecutiveCols = []string{
	"PERSON_NAME", "SEX", "AGE", "POSITION",
	"HIGH_DEGREE", "RESUME", "ORG_CODE",
}

var executiveExcel2Struct = map[string]string{
	"PERSON_NAME": "Name",
	"SEX":         "Sex",
	"AGE":         "Age",
	"POSITION":    "Position",
	"HIGH_DEGREE": "HighDegree",
	"RESUME":      "Resume",
	"ORG_CODE":    "OrgCode",
}

func ReadExecutive() (resInfos []dbmodels.Executive, err error) {
	return readExcel[dbmodels.Executive](
		"resources/所有高管.xlsx",
		usefulExecutiveCols, executiveExcel2Struct)
}

func readExcel[T dbmodels.Company | dbmodels.University | dbmodels.Executive](
	filePath string,
	usefulCols []string,
	excel2Struct map[string]string,
) (resInfos []T, err error) {

	var f *excelize.File

	// 读取数据
	if f, err = excelize.OpenFile(filePath); err != nil {
		logger.Error(err)
		return
	}
	defer func(f *excelize.File) {
		if err := f.Close(); err != nil {
			logger.Error(err)
		}
	}(f)

	// 获取 data 上所有单元格(第一个Sheet)
	// 突然发现有的是 data 有的是 sheet1，所以这里直接获取第一个sheet
	var rows [][]string
	sheets := f.GetSheetList()
	if len(sheets) != 1 {
		err = errors.New("sheet数量不为1")
		return
	}
	// 直接读第一个sheet
	if rows, err = f.GetRows(sheets[0]); err != nil {
		fmt.Println(err)
		return
	}

	// 获取有用行的索引
	usefulIndex := make([]int, 0)
	for index, titleCel := range rows[0] {
		if calc.IsTargetInArray[string](titleCel, usefulCols) {
			usefulIndex = append(usefulIndex, index)
		}
	}

	logger.Warning(usefulCols)
	logger.Warning(usefulIndex)

	// 保存数据
	resInfos = make([]T, 0)

	// 查询数据
	for _, row := range rows[1:] {

		// 反射创建结构体来记录每一行的数据
		// 直接 TypeOf(T) 无法通过编译，因此这里用 new(T) 来获取类型
		// 由于指定了多个泛型、无法确定结构体所占空间大小?
		// 参见 -> https://go.dev/doc/tutorial/generics
		resInfo := reflect.New(reflect.TypeOf(*new(T))).Elem()

		for j, colCell := range row {
			if !calc.IsTargetInArray[int](j, usefulIndex) {
				continue
			}

			if len(colCell) == 0 {
				err = errors.New("表格存在无效数据")
				return
			}

			// 通过反射将对应行的数据保存到结构体中
			field := resInfo.FieldByName(excel2Struct[rows[0][j]])

			switch field.Type() {
			case reflect.TypeOf(uint32(0)):
				num, _ := strconv.Atoi(colCell)
				field.SetUint(uint64(num))
			case reflect.TypeOf(""):
				field.SetString(colCell)
			case reflect.TypeOf(float32(0)):
				floatData, _ := strconv.ParseFloat(colCell, 32)
				field.SetFloat(floatData)
			default:
				logger.Error("未知类型" + field.Type().String())
			}

		}

		// 打印读取到的数据
		//logger.Info(resInfo.Interface())

		// 只保存这一行需要的数据，并添加到结果数组
		resInfos = append(resInfos, resInfo.Interface().(T))

	}
	return
}
