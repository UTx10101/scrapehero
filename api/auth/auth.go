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

func CreateToken(uname string, kid uint32, action string) (string, error) {
	claims := jwt.MapClaims{}
	claims["authorized"] = true
	
	if action == "AUTH" {
		claims["username"] = uname
	} else if action == "API" {
		claims["id"] = kid
	}
	
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	return token.SignedString([]byte(viper.GetString("api.secret_key")))
}

func ValidateToken(r *http.Request) error {
	tokenString := ExtractToken(r)
	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		//is token.Method type of/can be converted to *jwt.SigningMethodHMAC ?
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("Unexpected signing method: %v", token.Header["alg"])
		}
		return []byte(viper.GetString("api.secret_key")), nil
	})
	if err != nil {
		return err
	}
	if claims, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {
		Pretty(claims)
	}
	return nil
}

func ExtractToken(r *http.Request) string {
	keys := r.URL.Query()
	token := keys.Get("token")
	if token != "" {
		return token
	}
	bearerToken := r.Header.Get("Authorization")
	if len(strings.Split(bearerToken, " ")) == 2 {
		return strings.Split(bearerToken, " ")[1]
	}
	return ""
}

func ExtractTokenData(r *http.Request, action string) (string, uint32, error) {	
	tokenString := ExtractToken(r)
	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("Unexpected signing method: %v", token.Header["alg"])
		}
		return []byte(viper.GetString("api.secret_key")), nil
	})
	if err != nil {
		return 0, err
	}
	claims, ok := token.Claims.(jwt.MapClaims)
	if ok && token.Valid {
		uname, found := claims["username"]
		
		if kid, found := claims["id"]; found {
			id, err := strconv.ParseUint(fmt.Sprintf("%.0f", claims["id"]), 10, 32)
			if err != nil {
				return 0, err
			}
		}
		
		return string(uname), uint32(id), nil
	}
	return '',0, nil
}

// Pretty display the claims nicely in the terminal
func Pretty(data interface{}) {
	b, err := json.MarshalIndent(data, "", " ")
	if err != nil {
		log.Println(err)
		return
	}
	fmt.Println(string(b))
}