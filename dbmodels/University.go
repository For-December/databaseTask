package dbmodels

// University 985高校基本信息表
// UniversityID INT PRIMARY KEY AUTO_INCREMENT
// UniversityName VARCHAR(100)
// Province VARCHAR(50)
// City VARCHAR(50)
type University struct {
	UniversityID   uint32 `gorm:"not null;primaryKey;index"`
	UniversityName string `gorm:"not null;type:VARCHAR(255)"`
	Province       string `gorm:"not null;type:VARCHAR(50)"`
	City           string `gorm:"not null;type:VARCHAR(50)"`
}
