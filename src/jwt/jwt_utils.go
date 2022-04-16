package jwt

import (
	"crypto/rsa"
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"
	"net/http"
	"notes/src/config"
	"time"

	"github.com/golang-jwt/jwt"
)

type GoogleClaims struct {
	Email         string `json:"email"`
	EmailVerified bool   `json:"email_verified"`
	FirstName     string `json:"given_name"`
	LastName      string `json:"family_name"`
	jwt.StandardClaims
}

func getGooglePublicKey(keyID string) (*rsa.PublicKey, error) {
	response, err := http.Get("https://www.googleapis.com/oauth2/v1/certs")
	if err != nil {
		return nil, err
	}
	responseJson, err := ioutil.ReadAll(response.Body)
	if err != nil {
		return nil, err
	}
	unmarshalledJson := map[string]string{}
	err = json.Unmarshal(responseJson, &unmarshalledJson)
	if err != nil {
		return nil, err
	}
	privacyEnhancedMailKey, ok := unmarshalledJson[keyID]
	if !ok {
		return nil, errors.New("Google public key not found")
	}
	key, err := jwt.ParseRSAPublicKeyFromPEM([]byte(privacyEnhancedMailKey))
	return key, nil
}

func ValidateGoogleJWT(raw_token string) (GoogleClaims, error) {
	token, err := jwt.ParseWithClaims(
		raw_token,
		&GoogleClaims{},
		func(token *jwt.Token) (interface{}, error) {
            return getGooglePublicKey(fmt.Sprintf("%s", token.Header["kid"]))
        },
	)
    if err != nil {
        return GoogleClaims{}, err
    }
    claims, ok := token.Claims.(*GoogleClaims)
    if !ok {
        return GoogleClaims{}, errors.New("This JWT is not a Google JWT!")
    }
    if claims.Issuer != "accounts.google.com" && claims.Issuer != "https://accounts.google.com" {
        return GoogleClaims{}, errors.New("Invalid issuer (iss)!")
    }
    if claims.Audience != config.Config.Google.ClientID {
        return GoogleClaims{}, errors.New("Audience (aud) is invalid!")
    }
    if claims.ExpiresAt < time.Now().UTC().Unix() {
        return GoogleClaims{}, errors.New("JWT is expired!")
    }
    return *claims, nil
}
