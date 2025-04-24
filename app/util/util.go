package util

import (
	"errors"
	"fmt"
	"log"
	"os"
	"time"

	"github.com/golang-jwt/jwt/v5"
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
	return err==nil
}

const secretKey = "supersecret"

func GenerateToken(username string, id string) (string, error) {
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"username": username,
		"id": id,
		"expiration": time.Now().Add(time.Hour*2).Unix(),
	})
	return token.SignedString([]byte(secretKey))
}

func VerifyToken(token string) (string, error) {
	parsedToken, err := jwt.Parse(token, func(token *jwt.Token) (any, error) {
		_, ok := token.Method.(*jwt.SigningMethodHMAC)
		if !ok {
			return "", errors.New("unexpectd signing method")
		}
		return []byte(secretKey), nil
	})
	if err != nil {
		return "", err
	}

	if !parsedToken.Valid {
		return "", errors.New("invalid token")
	}

	_, ok := parsedToken.Claims.(jwt.MapClaims)
	if !ok {
		return "", errors.New("invalid token claims")
	}

	//id := claims["id"]
	return "", nil
}

func GetLocTimeZone() *time.Location {
	loc, err := time.LoadLocation(os.Getenv("TZ"))
	if err != nil {
		log.Fatalf("Error getting TZ - '%v': \n%v", os.Getenv("TZ"), err)
		return nil
	}
	fmt.Println()
	fmt.Println(loc)
	fmt.Println(loc.String())
	fmt.Println()
	return loc
}