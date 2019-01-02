package middleware

import (
	"os"
	"time"

	jwt "github.com/dgrijalva/jwt-go"
)

type MyClaims struct {
	Name string `json:"name"`
	jwt.StandardClaims
}

// Standared Claim properties
// Audience  string
// ExpiresAt int64
// Id        string
// IssuedAt  int64
// Issuer    string
// NotBefore int64
// Subject   string

func CreateJwtToken() (string, error) {

	jwtTokenSecret := []byte(getJwtTokenSecretKey())

	// Create the Claims
	claims := MyClaims{
		"Monir",
		jwt.StandardClaims{
			ExpiresAt: time.Now().Add(1 * time.Minute).Unix(),
			Issuer:    "Monir Hossain",
		},
	}

	rawToken := jwt.NewWithClaims(jwt.SigningMethodHS512, claims)
	token, err := rawToken.SignedString(jwtTokenSecret)

	if err != nil {
		return "", err
	}
	return token, nil
}

func PurseToken(t string) (string, error) {
	tokenString := t

	jwtTokenSecret := []byte(getJwtTokenSecretKey())

	var response string
	var errr error

	if tokenString == "" {
		return response, errr
	}

	token, err := jwt.ParseWithClaims(tokenString, &MyClaims{}, func(token *jwt.Token) (interface{}, error) {
		return []byte(jwtTokenSecret), nil
	})

	if claims, ok := token.Claims.(*MyClaims); ok {
		if time.Now().Unix() > claims.StandardClaims.ExpiresAt {
			response = "Your token is expired"
		} else if token.Valid {
			response = "Name: " + claims.Name
		}
	} else {
		errr = err
	}
	return response, errr
}

func getJwtTokenSecretKey() string {
	jwt_secret := os.Getenv("JWT_TOKEN_SECRET")

	if len(jwt_secret) > 0 {
		return jwt_secret
	}
	return ""
}
