package models

import (
	"notes/src/config"

	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

func getDB(databaseFileName string) (*gorm.DB, error) {
	db, err := gorm.Open(sqlite.Open(databaseFileName), &gorm.Config{
		Logger: logger.Default.LogMode(logger.Silent),
	})
	if err != nil {
		return nil, err
	}
	return db, nil
}

var DB *gorm.DB

func init() {
	var err error
	DB, err = getDB(config.Config.Database.Filename)
	if err != nil {
		panic(err)
	}
	DB.AutoMigrate(&User{}, &Note{})
}
