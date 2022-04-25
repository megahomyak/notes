package utils

import (
	"crypto/sha256"
	"encoding/base64"
	"errors"
	"fmt"
	"net/http"
	"notes/src/models"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

type AccessTokenNotFound struct {}

func (err *AccessTokenNotFound) Error() string {
	return "accessToken isn't provided!"
}

func GetAccessTokenHash(c *gin.Context) ([]byte, error) {
	encodedAccessToken, err := c.Cookie("access_token")
	if err != nil || encodedAccessToken == "" {
		return nil, &AccessTokenNotFound{}
	}
	decodedAccessToken, base64EncodingError := base64.StdEncoding.DecodeString(encodedAccessToken)
	if base64EncodingError != nil {
		return nil, base64EncodingError
	}
	accessTokenHash := sha256.Sum256(decodedAccessToken)
	return accessTokenHash[:], nil
}

func GetUserByToken(c *gin.Context) (*models.User, error) {
	token := models.AccessToken{}
	accessTokenHash, err := GetAccessTokenHash(c)
	if err != nil {
		return nil, err
	}
	userFindingError := models.DB.Where("hash = ?", accessTokenHash).Preload("Owner").Take(&token).Error
	if errors.Is(userFindingError, gorm.ErrRecordNotFound) {
		return nil, userFindingError
	} else if userFindingError == nil {
		token.ResetExpiration()
		return token.Owner, nil
	} else {
		c.Error(userFindingError)
		return nil, userFindingError
	}
}

func getNoteOr404(responseAdder func(c *gin.Context, noteID int64)) func(*gin.Context) *models.Note {
	return func(c *gin.Context) *models.Note {
		note := models.Note{}
		noteID := c.MustGet("note_id").(int64)
		noteGettingError := models.DB.Where("id = ?", noteID).Take(&note).Error
		if errors.Is(noteGettingError, gorm.ErrRecordNotFound) || noteGettingError != nil {
			if !errors.Is(noteGettingError, gorm.ErrRecordNotFound) {
				c.Error(noteGettingError)
			}
			responseAdder(c, noteID)
			return nil
		} else {
			return &note
		}
	}
}

var GetNoteOr404WithHTMLResponse func(c *gin.Context) *models.Note = getNoteOr404(
	func(c *gin.Context, noteID int64) {
		c.HTML(http.StatusNotFound, "note_not_found.tmpl", gin.H{"noteID": noteID})
	},
)

var GetNoteOr404WithJSONResponse func(c *gin.Context) *models.Note = getNoteOr404(
	func(c *gin.Context, noteID int64) {
		c.JSON(http.StatusNotFound, MakeJSONError(
			fmt.Sprintf("The note with the ID %d is not found!", noteID),
		))
	},
)
