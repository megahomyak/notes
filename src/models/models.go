package models

import (
	"time"
)

type User struct {
	ID         uint64
	FirstName  string  `gorm:"NOT NULL"`
	LastName   string  `gorm:"NOT NULL"`
	JWTSubject string  `gorm:"UniqueIndex;NOT NULL"`
	Notes      *[]Note `gorm:"ForeignKey:OwnerID"`
}

type AccessToken struct {
	Hash      []byte    `gorm:"primaryKey"`
	OwnerID   uint64    `gorm:"NOT NULL"`
	ExpiresIn time.Time `gorm:"NOT NULL"`
	Owner     *User     `gorm:"ForeignKey:OwnerID;Constraint:OnDelete:CASCADE"`
}

type Note struct {
	ID       uint64
	Name     string `gorm:"NOT NULL"`
	Contents string `gorm:"NOT NULL"`
	OwnerID  uint64 `gorm:"NOT NULL;UniqueIndex"`
	Owner    *User  `gorm:"ForeignKey:OwnerID;Constraint:OnDelete:CASCADE"`
}

var defaultTokenExpirationPeriod time.Duration = time.Hour * 24 * 30 * 6

func (token *AccessToken) ResetExpiration() {
	token.ExpiresIn = time.Now().Add(defaultTokenExpirationPeriod)
}
