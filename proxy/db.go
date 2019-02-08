package main

import (
	"gatecloud-boilerplate/proxy/configs"

	"github.com/jinzhu/gorm"
)

// InitDB creates the connection to database and returns a db handler
func InitDB() (*gorm.DB, error) {
	db, err := gorm.Open(configs.Configuration.DbEngine, configs.Configuration.DbConn)
	if err != nil {
		return nil, err
	}

	if configs.Configuration.AutoMigration {
		if err := db.AutoMigrate().Error; err != nil {
			return nil, err
		}
	}

	return db, nil
}
