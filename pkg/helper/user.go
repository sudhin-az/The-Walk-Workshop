package helper

import (
	"ecommerce_clean_architecture/pkg/utils/models"
	"time"

	"github.com/golang-jwt/jwt/v4"
	"golang.org/x/crypto/bcrypt"
)

type authCustomClaimsUsers struct {
	Id    int    `json:"id"`
	Email string `json:"email"`
	jwt.StandardClaims
}

func HashPassword(password string) (string, error) {
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		return "", err
	}
	return string(hashedPassword), nil
}

func GenerateTokenUsers(userID uint, userEmail string, expirationTime time.Time) (string, error) {

	claims := &authCustomClaimsUsers{
		Id:    int(userID), // Convert to int for use in claims
		Email: userEmail,
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: expirationTime.Unix(),
			IssuedAt:  time.Now().Unix(),
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	tokenString, err := token.SignedString([]byte("132457689"))

	if err != nil {
		return "", err
	}

	return tokenString, nil
}

func GenerateAccessToken(user models.UserDetailsResponse) (string, error) {

	expirationTime := time.Now().Add(15 * time.Minute)
	tokenString, err := GenerateTokenUsers(uint(user.Id), user.Email, expirationTime)
	if err != nil {
		return "", err
	}
	return tokenString, nil

}

func GenerateRefreshToken(user models.UserDetailsResponse) (string, error) {

	expirationTime := time.Now().Add(24 * 90 * time.Hour)
	tokeString, err := GenerateTokenUsers(uint(user.Id), user.Email, expirationTime)
	if err != nil {
		return "", err
	}
	return tokeString, nil

}
