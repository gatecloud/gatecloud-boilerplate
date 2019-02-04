package main

import (
	"gatecloud-boilerplate/proxy/configs"

	"github.com/jinzhu/gorm"
)

func InitDB() (*gorm.DB, error) {
	db, err := gorm.Open(configs.Configuration.Env.DbEngine, configs.Configuration.Env.DbConn)
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
