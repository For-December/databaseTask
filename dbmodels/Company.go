package dbmodels

// Company 上市公司基本信息
// CompanyCode VARCHAR(10) PRIMARY KEY 公司代码
// CompanyName VARCHAR(100) 公司名称
// RegisteredAddress VARCHAR(200) 注册地址
// EmployeeCount INT 员工人数
// RegisteredCapital DECIMAL(18, 2) 注册资金
// Industry VARCHAR(50) 行业
// StockExchange VARCHAR(50) 交易所
type Company struct {
	CompanyCode       string  `gorm:"not null;primaryKey;type:VARCHAR(10);index"`
	CompanyName       string  `gorm:"not null;type:VARCHAR(255)"`
	RegisteredAddress string  `gorm:"not null;type:VARCHAR(255)"`
	EmployeeCount     uint32  `gorm:"not null"`
	RegisteredCapital float32 `gorm:"not null;type:DECIMAL(20,8)"`
	Industry          string  `gorm:"not null;type:VARCHAR(50)"`
	StockExchange     string  `gorm:"not null;type:VARCHAR(50)"`

	// has many 关系
	Executive []Executive `gorm:"foreignKey:CompanyCode"`
}
