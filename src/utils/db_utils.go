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

const (
	WithNotes = true
	WithoutNotes = false
)

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

func GetUserByToken(c *gin.Context, withNotes bool) (*models.User, error) {
	token := models.AccessToken{}
	accessTokenHash, err := GetAccessTokenHash(c)
	if err != nil {
		return nil, err
	}
	query := models.DB.Where("hash = ?", accessTokenHash).Preload("Owner")
	if withNotes {
		query = query.Preload("Owner.Notes")
	}
	userFindingError := query.Take(&token).Error
	if errors.Is(userFindingError, gorm.ErrRecordNotFound) {
		return nil, userFindingError
	} else {
		token.ResetExpiration()
		return token.Owner, nil
	}
}

func getNoteOr404(responseAdder func(c *gin.Context, noteID int64)) func(*gin.Context) *models.Note {
	return func(c *gin.Context) *models.Note {
		note := models.Note{}
		noteID := c.MustGet("note_id").(int64)
		if errors.Is(
			models.DB.Where("id = ?", noteID).Take(&note).Error,
			gorm.ErrRecordNotFound,
		) {
			responseAdder(c, noteID)
			return nil
		}
		return &note
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
			fmt.Sprintf("The note with an ID %d is not found!", noteID),
		))
	},
)
