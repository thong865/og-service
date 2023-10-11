package auth

import (
	"errors"
	"fmt"
	"os"
	"time"

	jwt "github.com/dgrijalva/jwt-go"
	"github.com/joho/godotenv"
)

var refreshKey, secretKey []byte

type DataJson = map[string]interface{}

func init() {
	// get path from root dir
	err := godotenv.Load()
	rfKey := os.Getenv("REFRESH_KEY")
	scKey := os.Getenv("SECRET_KEY")
	if err != nil {
		panic(err)
	}
	refreshKey = []byte(rfKey)
	secretKey = []byte(scKey)
}

func GenarateAccessToken(data string) (string, error) {
	//  token is valid for 7days
	date := time.Now().Add(time.Minute * 7)
	// date := time.Now().Add(time.Second * 10)

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"mobile_no": data,
		"exp":       date.Unix(),
	})
	err := godotenv.Load()
	key := os.Getenv("SECRET_KEY")
	if err != nil {
		panic(err)
	}

	tokenString, error := token.SignedString([]byte(key))
	return tokenString, error
}

func GenarateRefreshToken(data DataJson) (string, error) {
	//  token is valid for 30days

	date := time.Now().Add(time.Hour * 24 * 30)

	// date := time.Now().Add(time.Second * 10)

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"authorized": true,
		"exp":        date.Unix(),
		"user":       data["mobile_no"],
	})
	err := godotenv.Load()
	key := os.Getenv("REFRESH_KEY")
	if err != nil {
		panic(err)
	}

	tokenString, error := token.SignedString([]byte(key))

	return tokenString, error
}

func ValidateRefreshToken(tokenString string) (DataJson, error) {
	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		// Don't forget to validate the alg is what you expect:
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("Unexpected signing method: %v", token.Header["alg"])
		}

		return refreshKey, nil
	})

	if err != nil {
		return DataJson{}, err
	}

	if !token.Valid {
		return DataJson{}, errors.New("invalid token")
	}

	return token.Claims.(jwt.MapClaims), nil
}

func ValidateAccessToken(tokenString string) (DataJson, error) {
	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		// Don't forget to validate the alg is what you expect:
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("Unexpected signing method: %v", token.Header["alg"])
		}

		return secretKey, nil
	})

	if err != nil {
		return DataJson{}, err
	}

	if !token.Valid {
		return DataJson{}, errors.New("invalid token")
	}

	return token.Claims.(jwt.MapClaims), nil
}
