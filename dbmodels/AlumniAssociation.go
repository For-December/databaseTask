package dbmodels

// AlumniAssociation 校友关联信息
// ExecutiveID INT
// UniversityCode VARCHAR(10)
// PRIMARY KEY (ExecutiveID, UniversityCode)
// FOREIGN KEY (ExecutiveID) REFERENCES Executive(ExecutiveID)
// FOREIGN KEY (UniversityCode) REFERENCES University(UniversityCode)
type AlumniAssociation struct {
	ExecutiveID    uint32 `gorm:"not null;primaryKey;index"`
	UniversityCode string `gorm:"not null;primaryKey;type:VARCHAR(10);index"`

	Executive  Executive  `gorm:"foreignKey:executive_id"`
	University University `gorm:"foreignKey:university_code"`
}
