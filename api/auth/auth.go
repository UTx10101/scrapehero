package auth

import (
	// buitin
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"strconv"
	"strings"

	// vendored
	"github.com/dgrijalva/jwt-go"
	"github.com/spf13/viper"
)

func CreateToken(uname string, kid string, action string) (string, error) {
	claims := jwt.MapClaims{}
	claims["authorized"] = true
	
	if action == "AUTH" {
		claims["username"] = uname
	} else if action == "API" {
		claims["id"] = kid
	}
	
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	return token.SignedString([]byte(viper.GetString("api.secret_key"))), nil
}

func ValidateToken(tokenString string) (string, string, error) {	
	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, nil, fmt.Errorf("Unexpected signing method: %v", token.Header["alg"])
		}
		return []byte(viper.GetString("api.secret_key")), nil
	})
	if err != nil {
		return nil, nil, err
	}
	if claims, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {
		if uname, found := claims["username"]; found {
			return string(uname), nil, nil
		}
		
		if kid, found := claims["id"]; found {
			return nil, string(kid), nil
		}
		
		return nil, nil, fmt.Errorf("malformed authtoken")
	}
	return nil, nil, fmt.Errorf("invalid authtoken")
}