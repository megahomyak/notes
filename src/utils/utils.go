package utils

import (
	"errors"
	"notes/src/models"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

const (
	WithNotes = true
	WithoutNotes = false
)

func GetUser(c *gin.Context, withNotes bool) (*models.User) {
	accessToken, accessTokenGettingError := c.Cookie("access_token")
	if accessTokenGettingError != nil {
		return nil
	} else {
		user := models.User{}
		query := models.DB.Where("access_token = ?", accessToken)
		if withNotes {
			query = query.Preload("Notes").Order("notes.id DESC")
		}
		if errors.Is(query.Take(&user).Error, gorm.ErrRecordNotFound) {
			return nil
		} else {
			return &user
		}
	}
}
