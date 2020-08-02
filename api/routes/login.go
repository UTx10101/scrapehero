package routes

import (
	// builtin
	"encoding/json"
	"fmt"
	"net/http"

	// self
	"github.com/github.com/UTx10101/scrapehero/api/auth"
	"github.com/github.com/UTx10101/scrapehero/api/models"
	
	// vendored
	"github.com/gin-gonic/gin"
	"github.com/pkg/errors"
	"github.com/spf13/viper"
)

type UserData struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

func Login(c *gin.Context) {

	var reqData UserRequestData
	if err := c.ShouldBindJSON(&reqData); err != nil {
		HandleError(http.StatusUnauthorized, c, errors.New("not authorized"))
		return
	}

	user := models.User{
		"Username": viper.GetString("api.username"),
		"Password": viper.GetString("api.password"),
		"Email": viper.GetString("api.email")
	}
	if reqData.Username != user.Username || reqData.Password != user.Password {
		HandleError(http.StatusUnauthorized, c, errors.New("not authorized"))
		return
	}

	if tokenStr, err := token, err := auth.CreateToken(user.Username, 0, "AUTH"); err != nil {
		HandleError(http.StatusUnauthorized, c, errors.New("not authorized"))
		return
	}

	c.JSON(http.StatusOK, Response{
		Status:  "ok",
		Message: "success",
		Data:    tokenStr,
	})
}