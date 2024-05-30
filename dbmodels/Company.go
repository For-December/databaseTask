package dbmodels

import "time"

// Company 上市公司基本信息
// CompanyCode VARCHAR(10) PRIMARY KEY
// CompanyName VARCHAR(100)
// RegisteredAddress VARCHAR(200)
// EmployeeCount INT
// RegisteredCapital DECIMAL(18, 2)
// Industry VARCHAR(50)
// StockExchange VARCHAR(50)
// EstablishedDate DATE
type Company struct {
	CompanyCode       string    `gorm:"not null;type:VARCHAR(10);primaryKey;index"`
	CompanyName       string    `gorm:"not null;type:VARCHAR(255)"`
	RegisteredAddress string    `gorm:"not null;type:VARCHAR(255)"`
	EmployeeCount     uint32    `gorm:"not null"`
	RegisteredCapital float32   `gorm:"not null;type:DECIMAL(18,2)"`
	Industry          string    `gorm:"not null;type:VARCHAR(50)"`
	StockExchange     string    `gorm:"not null;type:VARCHAR(50)"`
	EstablishedDate   time.Time `gorm:"not null;type:DATE"`
}
