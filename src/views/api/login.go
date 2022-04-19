package api

import (
	"crypto/sha256"
	"encoding/base64"
	"errors"
	"net/http"
	"notes/src/models"
	"notes/src/views/utils"
	"notes/src/views/utils/jwt"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func Login(c *gin.Context) {
	idToken := c.PostForm("id_token")
	if idToken == "" {
		c.JSON(http.StatusForbidden, utils.MakeJSONError("id_token wasn't provided!"))
		return
	}
	claims, err := jwt.ValidateGoogleJWT(idToken)
	if err != nil {
		c.JSON(http.StatusForbidden, utils.MakeJSONError(err.Error()))
		return
	}
	user := models.User{}
	var tokenContents []byte
	transactionError := models.DB.Transaction(func(tx *gorm.DB) error {
		if errors.Is(
			models.DB.Where("jwt_subject = ?", claims.Subject).Take(&user).Error,
			gorm.ErrRecordNotFound,
		) {
			user.FirstName = claims.FirstName
			user.LastName = claims.LastName
			user.JWTSubject = claims.Subject
			models.DB.Create(&user)
		}
		var accessTokenSlice []byte
		for {
			tokenContents = utils.MakeUniqueToken()
			accessToken := sha256.Sum256(tokenContents)
			accessTokenSlice = accessToken[:]
			var exists bool
			models.DB.Model(&models.AccessToken{}).Select("count(*) > 0").Where("hash = ?", accessTokenSlice).Find(&exists)
			// This PROBABLY introduces a race condition, if the second transaction got the same hash.
			// I don't know what's going to happen in this situation, especially in SQLite.
			// This is for the case when the same access token already exists.
			if !exists {
				break
			}
		}
		token := models.AccessToken{Owner: &user, Hash: accessTokenSlice}
		token.ResetExpiration()
		models.DB.Create(&token)
		return nil
	})
	if transactionError != nil {
		panic(transactionError)
	}
	utils.SetPermanentProtectedCookie(c, "access_token", base64.StdEncoding.EncodeToString(tokenContents))
	c.JSON(http.StatusOK, map[string]string{})
}
