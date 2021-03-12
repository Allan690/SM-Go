package model

import (
	"fmt"
	"github.com/dgrijalva/jwt-go"
	"os"
	"time"
)

type JWTService interface {
	GenerateToken(email string, userId string) (string, error)
	ValidateToken(token string)(*jwt.Token, error)
}

type authCustomClaims struct {
	Email string `json:"email"`
	ID string `json:"id"`
	jwt.StandardClaims
}

type jwtServices struct {
	secretKey string
	issuer string
}

// returns the jwtServices struct which implicitly implements the JWTServices interface
func JWTAuthService() JWTService {
	return &jwtServices{
		secretKey: getSecretKey(),
		issuer: os.Getenv("ISSUER"),
	}
}

func getSecretKey() string {
	secret := os.Getenv("SECRET")
	if secret == "" {
		secret = "secret"
	}
	return secret
}


// handles token generation
func (service *jwtServices) GenerateToken(email, userId string) (string, error) {
	claims :=&authCustomClaims{
		email,
		userId,
		jwt.StandardClaims{
			ExpiresAt: time.Now().Add(time.Hour * 48).Unix(),
			IssuedAt:  time.Now().Unix(),
			Issuer:  service.issuer,
		},
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	t, err := token.SignedString([]byte(service.secretKey))
	return t, err
}

// handles token validation
func (service *jwtServices) ValidateToken(encodedToken string)(*jwt.Token, error) {
	return jwt.Parse(encodedToken, func(token *jwt.Token) (interface{}, error) {
		if _, isValid := token.Method.(*jwt.SigningMethodHMAC); !isValid {
			return nil, fmt.Errorf("invalid token: %s", token.Header["alg"])
		}
		return []byte(service.secretKey), nil
	})
}

