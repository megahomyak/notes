package models

import "database/sql"

type User struct {
	ID          uint64
	FirstName   string
	LastName    string
	AccessToken sql.NullString `gorm:"uniqueIndex"`
	JWTSubject  string         `gorm:"uniqueIndex"`
	Notes       *[]Note
}

type Note struct {
	ID       uint64
	Name     string
	Contents string
	UserID   uint64
}
