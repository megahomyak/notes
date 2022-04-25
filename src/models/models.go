package models

import (
	"notes/src/logging"
	"time"
)

type User struct {
	ID         uint64
	FirstName  string         `gorm:"NOT NULL"`
	LastName   string         `gorm:"NOT NULL"`
	JWTSubject string         `gorm:"UniqueIndex;NOT NULL"`
}

func (user *User) GetNotes() []Note {
	var notes []Note
	if err := DB.Find(&notes, "owner_id = ?", user.ID).Error; err != nil {
		logging.LogError(err)
	}
	return notes
}

type AccessToken struct {
	Hash      []byte    `gorm:"primaryKey"`
	OwnerID   uint64    `gorm:"NOT NULL"`
	ExpiresIn time.Time `gorm:"NOT NULL"`
	Owner     *User     `gorm:"ForeignKey:OwnerID;Constraint:OnDelete:CASCADE"`
}

var defaultTokenExpirationPeriod time.Duration = time.Hour * 24 * 30 * 6

func (token *AccessToken) ResetExpiration() {
	if err := DB.Model(&AccessToken{}).Where("hash = ?", token.Hash).Update("expires_in", time.Now().Add(defaultTokenExpirationPeriod)).Error; err != nil {
		logging.LogError(err)
	}
}

type Note struct {
	ID       uint64
	Name     string `gorm:"NOT NULL"`
	Contents string `gorm:"NOT NULL"`
	OwnerID  uint64 `gorm:"NOT NULL;Index"`
	Owner    *User  `gorm:"ForeignKey:OwnerID;Constraint:OnDelete:CASCADE"`
}
