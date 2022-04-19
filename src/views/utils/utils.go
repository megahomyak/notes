package utils

import (
	"errors"
	"notes/src/models"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
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

func SetPermanentProtectedCookie(c *gin.Context, cookieName string, cookieContents string) {
	c.SetCookie(cookieName, cookieContents, 2147483647, "/", "", true, true)
}

func AddCSRFToken(c *gin.Context, templateData map[string]interface{}) {
	csrfToken := uuid.New().String()
	SetPermanentProtectedCookie(c, "csrf_token", csrfToken)
	templateData["csrfToken"] = csrfToken
}
