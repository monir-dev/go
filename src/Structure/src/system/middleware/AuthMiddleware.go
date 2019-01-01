package middleware

import (
	"os"
	"time"

	jwt "github.com/dgrijalva/jwt-go"
)

type MyClaims struct {
	Foo string `json:"foo"`
	jwt.StandardClaims
}

func CreateJwtToken() (string, error) {

	jwtTokenSecret := []byte(getJwtTokenSecretKey())

	// Create the Claims
	claims := MyClaims{
		"Monir",
		jwt.StandardClaims{
			ExpiresAt: time.Now().Add(24 * time.Hour).Unix(),
			Issuer:    "test",
		},
	}

	rawToken := jwt.NewWithClaims(jwt.SigningMethodHS512, claims)
	token, err := rawToken.SignedString(jwtTokenSecret)

	if err != nil {
		return "", err
	}
	return token, nil
}

func PurseToken(t string) (string, string, error) {
	tokenString := t

	jwtTokenSecret := []byte(getJwtTokenSecretKey())

	var response string
	var errorMsg string
	var errr error

	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		return []byte(jwtTokenSecret), nil
	})

	if token.Valid {
		response = "You look nice today"
	} else if ve, ok := err.(*jwt.ValidationError); ok {
		if ve.Errors&jwt.ValidationErrorMalformed != 0 {
			errorMsg = "That's not even a token"
		} else if ve.Errors&(jwt.ValidationErrorExpired|jwt.ValidationErrorNotValidYet) != 0 {
			errorMsg = "Token is either expired or not active yet"
		} else {
			errorMsg = "Couldn't handle this token"
			errr = err
		}
	} else {
		errorMsg = "Couldn't handle this token"
		errr = err
	}

	return response, errorMsg, errr
}

func getJwtTokenSecretKey() string {
	jwt_secret := os.Getenv("JWT_TOKEN_SECRET")

	if len(jwt_secret) > 0 {
		return jwt_secret
	}
	return ""
}
