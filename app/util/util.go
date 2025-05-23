package util

import (
	"errors"
	"log"
	"time"

	"github.com/golang-jwt/jwt/v5"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"golang.org/x/crypto/bcrypt"
)

// About password
func ValidatePassword(pass, confirmPass string) bool {
	return pass == confirmPass
}

func HashPassword(pass string) (string, error) {
	bytes, err := bcrypt.GenerateFromPassword([]byte(pass), 14)
	return string(bytes), err
}

func CheckPasswordHash(pass, hashedPass string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(hashedPass), []byte(pass))
	return err == nil
}

const secretKey = "supersecret"

func GenerateToken(email string, id primitive.ObjectID) (string, error) { // Use ObjectID for ID
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"email":      email,
		"id":         id.Hex(), // Convert ObjectID to Hex string
		"expiration": time.Now().Add(2 * time.Hour).Unix(),
	})
	return token.SignedString([]byte(secretKey))
}

func ValidateToken(token string) (map[string]any, error) {
	parsedToken, err := jwt.Parse(token, func(token *jwt.Token) (any, error) {
		_, ok := token.Method.(*jwt.SigningMethodHMAC)
		if !ok {
			return nil, errors.New("unexpected signing method")
		}
		return []byte(secretKey), nil
	})
	if err != nil {
		return nil, err
	}

	if !parsedToken.Valid {
		return nil, errors.New("invalid token")
	}

	claims, ok := parsedToken.Claims.(jwt.MapClaims)
	if !ok {
		return nil, errors.New("invalid token claims")
	}

	return claims, nil
}

const location = "America/Sao_Paulo"

func GetLocTimeZone() *time.Location {
	loc, err := time.LoadLocation(location)
	if err != nil {
		log.Fatalf("Error getting TZ - '%v': \n%v", location, err)
		return nil
	}
	return loc
}