package api

import (
	"database/sql"
	"errors"
	"net/http"
	"notes/src/jwt"
	"notes/src/models"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"gorm.io/gorm"
)

type loginRequest struct {
	IDToken string `json:"id_token"`
}

func Login(c *gin.Context) {
	var loginRequest loginRequest
	if c.BindJSON(&loginRequest) != nil {
		return
	}
	claims, err := jwt.ValidateGoogleJWT(loginRequest.IDToken)
	if err != nil {
		c.JSON(http.StatusForbidden, map[string]string{"error": err.Error()})
		return
	}
	user := models.User{}
	var accessToken string
	creationFailed := errors.Is(
		models.DB.Where("jwt_subject = ?", claims.Subject).Take(&user).Error,
		gorm.ErrRecordNotFound,
	)
	if creationFailed || !user.AccessToken.Valid {
		if creationFailed {
			user.FirstName = claims.FirstName
			user.LastName = claims.LastName
			user.JWTSubject = claims.Subject
		}
		for {
			accessToken = uuid.New().String()
			user.AccessToken = sql.NullString{String: accessToken, Valid: true}
			if err := models.DB.Create(&user).Error; errors.Is(err, gorm.ErrInvalidTransaction) {
				// Probably, UUID was matching with the other UUID.
				// I know that the probability of that is very low, but I'll handle it nonetheless.
				// I don't know how to generate completely unique UUIDs, so we're just gonna
				// hope that the next one is gonna be disengaged.
				// Or the one after it. Or the one after after it. Or the one after after after it...
				continue
			} else {
				break
			}
		}
	} else {
		accessToken = user.AccessToken.String
	}
	c.SetCookie("access_token", accessToken, 2147483647, "/", "", true, true)
	c.JSON(http.StatusOK, map[string]string{})
}
