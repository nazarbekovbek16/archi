package handlers

import (
	"errors"
	"fmt"
	"github.com/golang-jwt/jwt/v4"
	"github.com/labstack/echo/v4"
	"log"
	"net/http"
	"strconv"
	"time"
)

type envelope map[string]any

func WriteCookie(c echo.Context, v string) error {
	cookie := new(http.Cookie)

	cookie.Name = "token"
	cookie.Value = v
	loc, _ := time.LoadLocation("Asia/Almaty")
	cookie.Expires = time.Now().In(loc).Add(24 * time.Hour)

	c.SetCookie(cookie)

	return nil
}
func ReadCookie(c echo.Context) (string, error) {
	cookie, err := c.Cookie("token")
	if err != nil {
		return "", err
	}
	return cookie.Value, nil
}
func DeleteCookie(c echo.Context) error {
	cookie := new(http.Cookie)
	cookie.Name = "token"
	loc, _ := time.LoadLocation("Asia/Almaty")
	cookie.Expires = time.Now().In(loc)

	c.SetCookie(cookie)
	return nil
}

func ValidateJWT(secretKey string, tokenString string) (*jwt.Token, error) {
	return jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("Unexpected signing method: %v", token.Header["alg"])
		}
		return []byte(secretKey), nil
	})
}
func ReadIDParam(c echo.Context) (int, error) {
	params := c.Param("id")
	id, err := strconv.Atoi(params)
	if err != nil || id < 1 {
		return 0, errors.New("invalid id parameter")
	}
	return id, nil
}

func ExtractClaims(secretKey string, tokenStr string) (jwt.MapClaims, error) {
	hmacSecretString := secretKey
	hmacSecret := []byte(hmacSecretString)
	token, err := jwt.Parse(tokenStr, func(token *jwt.Token) (interface{}, error) {
		return hmacSecret, nil
	})

	if err != nil {
		return nil, err
	}

	if claims, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {
		return claims, nil
	} else {
		log.Printf("Invalid JWT Token")
		return nil, err
	}
}
