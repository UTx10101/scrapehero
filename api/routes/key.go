package routes

import (
	// builtin
	"net/http"
	"options"
	
	// self
	"github.com/UTx10101/scrapehero/constants"
	"github.com/UTx10101/scrapehero/db"
	"github.com/UTx10101/scrapeher/models"
	
	// vendored
	"github.com/gin-gonic/gin"
	"github.com/globalsign/mgo/bson"
)

func GetAPIKeys(c *gin.Context) {
	query := bson.M{}
	
	if keys, err := models.GetKeys(query); err != nil {
		HandleError(http.StatusInternalServerError, c, err)
		return
	}

	if total, err := models.GetKeysCount(query); err != nil {
		HandleError(http.StatusInternalServerError, c, err)
		return
	}

	c.JSON(http.StatusOK, ListResponse{
		Status:  "ok",
		Message: "success",
		Data:    keys,
		Total:   total,
	})
}

func CreateAPIKey(c *gin.Context) {
	var ak models.APIKey
	
	ak.Status = "active"
	
	if err := ak.Create(); err != nil {
		HandleError(http.StatusInternalServerError, c, err)
		return
	}

	c.JSON(http.StatusOK, Response{
		Status:  "ok",
		Message: "success",
	})
}

func ModAPIKeyStatus(c *gin.Context) {
	kid := c.Param("kid")
	status := c.Param("status")

	if !bson.IsObjectIdHex(pid) {
		HandleErrorF(http.StatusBadRequest, c, "invalid key id")
		return
	}
	
	if status != "active" && status != "inactive" {
		HandleErrorF(http.StatusBadRequest, c, "invalid status for a key")
		return
	}

	if err := models.UpdateKeyStatus(bson.ObjectIdHex(kid, status)); err != nil {
		HandleError(http.StatusInternalServerError, c, err)
		return
	}

	c.JSON(http.StatusOK, Response{
		Status:  "ok",
		Message: "success",
	})
}

func DeleteAPIKey(c *gin.Context) {
	kid := c.Param("kid")

	if !bson.IsObjectIdHex(kid) {
		HandleErrorF(http.StatusBadRequest, c, "invalid key id")
		return
	}

	if err := models.RemoveKey(bson.ObjectIdHex(kid)); err != nil {
		HandleError(http.StatusInternalServerError, c, err)
		return
	}

	c.JSON(http.StatusOK, Response{
		Status:  "ok",
		Message: "success",
	})
}