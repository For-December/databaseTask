package dbmodels

// University 985高校基本信息表
// UniversityCode VARCHAR(10) PRIMARY KEY
// UniversityName VARCHAR(100)
// Province VARCHAR(50)
// City VARCHAR(50)
type University struct {
	UniversityCode string `gorm:"not null;type:VARCHAR(10);primaryKey;index"`
	UniversityName string `gorm:"not null;type:VARCHAR(255)"`
	Province       string `gorm:"not null;type:VARCHAR(50)"`
	City           string `gorm:"not null;type:VARCHAR(50)"`
}
