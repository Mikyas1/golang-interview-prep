package models

import (
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

func CreateSqliteDb(name string) (*gorm.DB, error) {
	db, err := gorm.Open(sqlite.Open(name), &gorm.Config{})
	if err != nil {
		return nil, err
	}
	return db, nil
}

func MigrateToSqliteDb(db *gorm.DB, models []interface{}) error {
	for _, model := range models {
		err := db.AutoMigrate(model)
		if err != nil {
			return err
		}
	}
	return nil
}
