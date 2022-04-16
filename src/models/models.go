package models

import "database/sql"

type User struct {
	ID          uint64 `gorm:"primarykey"`
	FirstName   string
	LastName    string
	AccessToken sql.NullString `gorm:"uniqueIndex"`
}

type Note struct {
	ID       uint64
	Contents string
	UserID   uint64
	User     User `gorm:"foreignKey:UserID"`
}
