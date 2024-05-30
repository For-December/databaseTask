package dbmodels

// AlumniAssociation 校友关联信息
// ExecutiveID INT
// UniversityID INT
// PRIMARY KEY (ExecutiveID, UniversityID)
// FOREIGN KEY (ExecutiveID) REFERENCES Executive(ExecutiveID)
// FOREIGN KEY (UniversityID) REFERENCES University(UniversityID)
type AlumniAssociation struct {
	ExecutiveID  uint32 `gorm:"not null;primaryKey;index"`
	UniversityID uint32 `gorm:"not null;primaryKey;index"`

	Executive  Executive  `gorm:"foreignKey:ExecutiveID"`
	University University `gorm:"foreignKey:UniversityID"`
}
