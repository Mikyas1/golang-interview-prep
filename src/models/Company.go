package models

import "gorm.io/gorm"

type Company struct {
	gorm.Model
	Name string `json:"name" gorm:"unique;not null"`
}

func CreateCompanies(db *gorm.DB) error {
	company := []Company{
		{
			Name: "ABC Company",
		},
		{
			Name: "DEF Company",
		},
	}

	result := db.Create(&company)
	if result.Error != nil {
		return result.Error
	}
	return nil
}
