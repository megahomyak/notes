package views

import (
	"database/sql"
	"net/http"
	jwt "notes/src/jwt"
	"notes/src/models"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

type LoginRequest struct {
	Token string `json:"token"`
}

func Login(c *gin.Context) {
	var loginRequest LoginRequest
	if c.BindJSON(&loginRequest) != nil {
		return
	}
	claims, err := jwt.ValidateGoogleJWT(loginRequest.Token)
	if err != nil {
		c.JSON(http.StatusForbidden, map[string]string{"error": err.Error()})
		return
	}
	var accessToken string
	for {
		accessToken = uuid.New().String()
		newUser := &models.User{
			FirstName: claims.FirstName,
			LastName: claims.LastName,
			AccessToken: sql.NullString{String: accessToken, Valid: true},
		}
		if err := models.DB.Create(newUser).Error; err != nil {
			// Probably, UUID was matching with the other UUID.
			// I know that probability of that is very low, but I'll handle it nonetheless.
			// I don't know how to generate completely unique UUIDs, so we're just gonna
			// hope that the next one is gonna be disengaged.
			// Or the one after it. Or the one after after it. Or the one after after after it...
			continue
		} else {
			break
		}
	}
	c.JSON(http.StatusOK, map[string]string{"access_token": accessToken})
}
