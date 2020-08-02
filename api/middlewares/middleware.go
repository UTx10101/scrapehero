package middlewares

import (
	// builtin
	"net/http"
	"strings"
	
	// self
	"github.com/UTx10101/scrapehero/api/auth"
	"github.com/UTx10101/scrapehero/api/routes"
	"github.com/UTx10101/scrapehero/api/models"

	// vendored
	"github.com/gin-gonic/gin"
	"github.com/spf13/viper"
)

func AuthMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		tokenStr := c.GetHeader("Authorization")

		if user, _, err := auth.ValidateToken(tokenStr); err != nil || user != viper.GetString("api.username") {
			c.AbortWithStatusJSON(http.StatusUnauthorized, routes.Response{
				Status:  "ok",
				Message: "unauthorized",
				Error:   "unauthorized",
			})
			return
		}

		c.Next()
	}
}

func APIAuthMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		tokenStr := c.GetHeader("Authorization")

		if err := models.CheckAPIKey(tokenStr); err != nil {
			c.AbortWithStatusJSON(http.StatusUnauthorized, routes.Response{
				Status:  "ok",
				Message: "unauthorized",
				Error:   err.Error(),
			})
			return
		}

		c.Next()
	}
}

// This enables us interact with the React Frontend
func CORSMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.Writer.Header().Set("Access-Control-Allow-Origin", "*")
		c.Writer.Header().Set("Access-Control-Allow-Credentials", "true")
		c.Writer.Header().Set("Access-Control-Allow-Headers", "Content-Type, Content-Length, Accept-Encoding, X-CSRF-Token, Authorization, accept, origin, Cache-Control, X-Requested-With")
		c.Writer.Header().Set("Access-Control-Allow-Methods", "POST, OPTIONS, GET, PUT, PATCH, DELETE")

		if c.Request.Method == "OPTIONS" {
			c.AbortWithStatus(204)
			return
		}
		c.Next()
	}
}