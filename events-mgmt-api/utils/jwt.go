package utils

import (
	"errors"
	"time"

	"github.com/golang-jwt/jwt/v5"
)

const secretKey = "momomomo"

func GenerateToken(email string, uid int64) (string, error) {
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"email": email,
		"uid":   uid,
		"exp":   time.Now().Add(time.Hour * 2).Unix(),
	})
	return token.SignedString([]byte(secretKey))
}

func VerifyToken(tokenStr string) (int64, error) {

	token, err := jwt.Parse(tokenStr, func(t *jwt.Token) (interface{}, error) {
		if _, ok := t.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, errors.New("unsupported signing method")
		}
		return []byte(secretKey), nil
	})

	if err != nil {
		return 0, errors.New("could not parse token")
	}

	if !token.Valid {
		return 0, errors.New("invalid token")
	}

	claim, ok := token.Claims.(jwt.MapClaims)
	if !ok {
		return 0, errors.New("invalid token claim")
	}

	// email := claim["email"].(string)
	uid := int64(claim["uid"].(float64))
	return uid, nil
}
