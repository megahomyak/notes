package api

import (
	"crypto/sha256"
	"database/sql"
	"encoding/base64"
	"errors"
	"net/http"
	"notes/src/models"
	"notes/src/utils"
	"notes/src/utils/jwt"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func SignIn(c *gin.Context) {
	idToken := c.PostForm("id_token")
	if idToken == "" {
		c.JSON(http.StatusUnauthorized, utils.MakeJSONError("id_token wasn't provided!"))
		return
	}
	claims, err := jwt.ValidateGoogleJWT(idToken)
	if err != nil {
		c.JSON(http.StatusUnauthorized, utils.MakeJSONError(err.Error()))
		return
	}
	user := models.User{}
	var tokenContents []byte
	transactionOptions := sql.TxOptions{Isolation: sql.LevelSerializable}
	transactionError := models.DB.Transaction(func(tx *gorm.DB) error {
		userFindingError := models.DB.Where("jwt_subject = ?", claims.Subject).Take(&user).Error;
		if errors.Is(userFindingError, gorm.ErrRecordNotFound) {
			user.FirstName = claims.FirstName
			user.LastName = claims.LastName
			user.JWTSubject = claims.Subject
			if err := models.DB.Create(&user).Error; err != nil {
				c.Error(err)
			}
		} else if userFindingError != nil {
			c.Error(userFindingError)
		}
		var accessTokenSlice []byte
		for {
			tokenContents = utils.MakeUniqueToken()
			accessToken := sha256.Sum256(tokenContents)
			accessTokenSlice = accessToken[:]
			var exists bool
			if err := models.DB.Model(&models.AccessToken{}).Select("count(*) > 0").Where("hash = ?", accessTokenSlice).Find(&exists).Error; err != nil {
				c.Error(err)
			}
			if !exists {
				break
			}
		}
		token := models.AccessToken{Owner: &user, Hash: accessTokenSlice}
		token.ResetExpiration()
		if err := models.DB.Create(&token).Error; err != nil {
			c.Error(err)
		}
		return nil
	}, &transactionOptions)
	if transactionError != nil {
		c.Error(transactionError)
	}
	utils.SetPermanentProtectedCookie(c, "access_token", base64.StdEncoding.EncodeToString(tokenContents))
	c.JSON(http.StatusOK, gin.H{})
}
