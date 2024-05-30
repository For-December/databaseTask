package dbmodels

// Executive 上市公司高管基本信息
// ExecutiveID INT PRIMARY KEY AUTO_INCREMENT
// Name VARCHAR(50)
// Gender VARCHAR(10)
// Age INT
// Position VARCHAR(50)
// Resume TEXT
// CompanyCode VARCHAR(10)
// FOREIGN KEY (CompanyCode) REFERENCES Company(CompanyCode)
type Executive struct {
	ExecutiveID uint32 `gorm:"not null;primaryKey;index;autoIncrement"`
	Name        string `gorm:"not null;type:VARCHAR(50)"`
	Gender      string `gorm:"not null;type:VARCHAR(10)"`
	Age         uint32 `gorm:"not null"`
	Position    string `gorm:"not null;type:VARCHAR(50)"`
	Resume      string `gorm:"not null;type:TEXT"`
	CompanyCode string `gorm:"not null;index;type:VARCHAR(10)"`
}
