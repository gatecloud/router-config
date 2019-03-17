package main

import (
	"router-config/configs"
	"router-config/models"

	"github.com/jinzhu/gorm"
)

// InitDB initialises the database
func InitDB() (*gorm.DB, error) {
	db, err := gorm.Open(configs.Configuration.DbEngine, configs.Configuration.DbConn)
	if err != nil {
		return nil, err
	}

	if !configs.Configuration.Production {
		if err = db.AutoMigrate(
			models.Project{},
			models.Template{},
		).Error; err != nil {
			return nil, err
		}
	}

	return db, nil
}
