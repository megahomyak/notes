package models

import (
	"notes/src/config"

	"github.com/gin-gonic/gin"
	"gorm.io/driver/postgres"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

var DatabaseNameToDialectorCreator = map[string]func(string) gorm.Dialector{
	"postgres": postgres.Open,
	"sqlite": sqlite.Open,
}

func getDB() (*gorm.DB, error) {
	var logMode logger.LogLevel
	if gin.Mode() == gin.ReleaseMode {
		logMode = logger.Silent
	} else {
		logMode = logger.Info
	}
	backendName := config.Config.Database.DefaultBackend
	dialector := DatabaseNameToDialectorCreator[backendName](
		config.Config.Database.Arguments[backendName],
	)
	db, err := gorm.Open(dialector, &gorm.Config{
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
	DB, err = getDB()
	if err != nil {
		panic(err)
	}
	DB.AutoMigrate(&User{}, &Note{}, &AccessToken{})
}
