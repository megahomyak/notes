package models

import "database/sql"

type User struct {
	ID          uint64
	FirstName   string         `gorm:"not null"`
	LastName    string         `gorm:"not null"`
	AccessToken sql.NullString `gorm:"uniqueIndex;not null"`
	JWTSubject  string         `gorm:"uniqueIndex;not null"`
	Notes       *[]Note
}

type Note struct {
	ID       uint64
	Name     string `gorm:"not null"`
	Contents string `gorm:"not null"`
	UserID   uint64 `gorm:"not null"`
	Owner    *User  `gorm:"ForeignKey:UserID"`
}
