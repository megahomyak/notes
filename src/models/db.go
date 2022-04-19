package models

import (
	"notes/src/config"

	"github.com/gin-gonic/gin"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

func getDB(databaseFileName string) (*gorm.DB, error) {
	var logMode logger.LogLevel
	if gin.Mode() == gin.ReleaseMode {
		logMode = logger.Silent
	} else {
		logMode = logger.Info
	}
	db, err := gorm.Open(sqlite.Open(databaseFileName), &gorm.Config{
		Logger: logger.Default.LogMode(logMode),
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
	DB.AutoMigrate(&User{}, &Note{}, &AccessToken{})
}
