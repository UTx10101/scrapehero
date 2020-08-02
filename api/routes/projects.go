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

func GetProjects(c *gin.Context) {
	query := bson.M{}
	
	if projects, err := models.GetProjects(query); err != nil {
		HandleError(http.StatusInternalServerError, c, err)
		return
	}

	if total, err := models.GetProjectsCount(query); err != nil {
		HandleError(http.StatusInternalServerError, c, err)
		return
	}

	c.JSON(http.StatusOK, ListResponse{
		Status:  "ok",
		Message: "success",
		Data:    projects,
		Total:   total,
	})
}

func GetProject(c *gin.Context) {
	pid := c.Param("pid")
	
	if !bson.IsObjectIdHex(pid) {
		HandleErrorF(http.StatusBadRequest, c, "invalid project id")
		return
	}
	
	if project, err := models.GetProject(bson.ObjectIdHex(pid)); err != nil {
		HandleError(http.StatusInternalServerError, c, err)
		return
	}

	c.JSON(http.StatusOK, Response{
		Status:  "ok",
		Message: "success",
		Data:    project,
	})
}

func EditProject(c *gin.Context) {
	pid := c.Param("pid")
	action := c.Param("action")
	
	if !bson.IsObjectIdHex(pid) {
		HandleErrorF(http.StatusBadRequest, c, "invalid project id")
		return
	}
	
	if action == "1" {
		if err := utils.StartEditor(); err != nil {
			HandleError(http.StatusInternalServerError, c, err)
			return
		}
	} else if action == "0" {
		if err := utils.StopEditor(); err != nil {
			HandleError(http.StatusInternalServerError, c, err)
			return
		}
	} else {
		HandleErrorF(http.StatusBadRequest, c, "invalid editor action")
		return
	}
	
	c.JSON(http.StatusOK, Response{
		Status:  "ok",
		Message: "success",
	})
}

func CreateProject(c *gin.Context) {
	var p models.Project
	
	if err := c.ShouldBindJSON(&p); err != nil {
		HandleError(http.StatusBadRequest, c, err)
		return
	}
	
	if err := p.Create(); err != nil {
		HandleError(http.StatusInternalServerError, c, err)
		return
	}

	c.JSON(http.StatusOK, Response{
		Status:  "ok",
		Message: "success",
	})
}

func ModProject(c *gin.Context) {
	pid := c.Param("pid")

	if !bson.IsObjectIdHex(pid) {
		HandleErrorF(http.StatusBadRequest, c, "invalid project id")
		return
	}

	var prj models.Project
	if err := c.ShouldBindJSON(&prj); err != nil {
		HandleError(http.StatusBadRequest, c, err)
		return
	}

	if err := models.UpdateProject(bson.ObjectIdHex(pid), prj); err != nil {
		HandleError(http.StatusInternalServerError, c, err)
		return
	}

	c.JSON(http.StatusOK, Response{
		Status:  "ok",
		Message: "success",
	})
}

func DeleteProject(c *gin.Context) {
	pid := c.Param("pid")

	if !bson.IsObjectIdHex(pid) {
		HandleErrorF(http.StatusBadRequest, c, "invalid project id")
		return
	}

	if err := models.RemoveProject(bson.ObjectIdHex(pid)); err != nil {
		HandleError(http.StatusInternalServerError, c, err)
		return
	}

	c.JSON(http.StatusOK, Response{
		Status:  "ok",
		Message: "success",
	})
}