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

func GetUser(c *gin.Context, withNotes bool) (*models.User, error) {
	accessToken, accessTokenGettingError := c.Cookie("access_token")
	if accessTokenGettingError != nil {
		return nil, errors.New("access_token wasn't found!")
	} else {
		user := models.User{}
		query := models.DB.Where("access_token = ?", accessToken)
		if withNotes {
			query = query.Preload("Notes").Order("id DESC")
		}
		if errors.Is(query.Take(&user).Error, gorm.ErrRecordNotFound) {
			return nil, errors.New(
				"User with the provided access_token wasn't found in the database",
			)
		} else {
			return &user, nil
		}
	}
}

func MakeJSONError(error_text string) map[string]string {
	return map[string]string{"error": error_text}
}
