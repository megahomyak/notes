package utils

import (
	"crypto/sha256"
	"encoding/base64"
	"encoding/hex"
	"errors"
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
	hexHash := hex.EncodeToString(accessTokenHash)
	query := models.DB.Where("lower(hex(hash)) = ?", hexHash).Preload("Owner")  // A nasty workaround.
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
