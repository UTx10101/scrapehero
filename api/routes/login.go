package routes

import (
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/github.com/UTx10101/scrapehero/api/auth"
	"github.com/github.com/UTx10101/scrapehero/api/models"
	"github.com/github.com/UTx10101/scrapehero/api/security"
	"github.com/github.com/UTx10101/scrapehero/api/utils"
	"golang.org/x/crypto/bcrypt"
	"github.com/spf13/viper"
)

func Login(c *gin.Context) {

	//clear previous error if any
	errList = map[string]string{}

	body, err := ioutil.ReadAll(c.Request.Body)
	if err != nil {
		c.JSON(http.StatusUnprocessableEntity, gin.H{
			"status":      http.StatusUnprocessableEntity,
			"first error": "Unable to get request",
		})
		return
	}
	user := models.User{}
	err = json.Unmarshal(body, &user)
	if err != nil {
		c.JSON(http.StatusUnprocessableEntity, gin.H{
			"status": http.StatusUnprocessableEntity,
			"error":  "Cannot unmarshal body",
		})
		return
	}
	user.Prepare()
	errorMessages := user.Validate()
	if len(errorMessages) > 0 {
		c.JSON(http.StatusUnprocessableEntity, gin.H{
			"status": http.StatusUnprocessableEntity,
			"error":  errorMessages,
		})
		return
	}
	userData, err := SignIn(user.Username, user.Password)
	if err != nil {
		formattedError := utils.FormatError(err.Error())
		c.JSON(http.StatusUnprocessableEntity, gin.H{
			"status": http.StatusUnprocessableEntity,
			"error":  formattedError,
		})
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"status":   http.StatusOK,
		"response": userData,
	})
}

func SignIn(username, password string) (map[string]interface{}, error) {

	var err error

	userData := make(map[string]interface{})

	user := models.User{
		"Username": viper.GetString("api.username"),
		"Password": viper.GetString("api.password"),
		"Email": viper.GetString("api.email")
	}
	
	if username == user.Username {
		err = errors.New("error getting the user: user not found")
		return nil, err
	}
	
	if password == user.Password {
		err = errors.New("error checking the password: incorrect password")
		return nil, err
	}
	token, err := auth.CreateToken(user.Username, 0, "AUTH")
	if err != nil {
		return nil, err
	}
	userData["token"] = token
	userData["username"] = user.Username
	userData["email"] = user.Email
	
	return userData, nil
}