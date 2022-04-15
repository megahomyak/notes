package main

import (
	"database/sql"
	"notes/src/config"

	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

type User struct {
	ID        int64
	FirstName string
	LastName  string
	Token     sql.NullString `gorm:"uniqueIndex"`
}

type Note struct {
	ID       int64
	Contents string
	UserID   int64
	User     `gorm:"foreignKey:UserID"`
}

func getDB(databaseFileName string) (*gorm.DB, error) {
	db, err := gorm.Open(sqlite.Open(databaseFileName), &gorm.Config{})
	if err != nil {
		return nil, err
	}
	return db, nil
}

var db *gorm.DB

func init() {
	var err error
	db, err = getDB(config.Config.Database.Filename)
	if err != nil {
		panic(err)
	}
	db.AutoMigrate(&User{}, &Note{})
}
