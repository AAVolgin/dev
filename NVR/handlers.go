package main

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
)

func ListCameras(c *gin.Context) {
	var cams []Camera
	DB.Find(&cams)
	c.JSON(http.StatusOK, cams)
}

func AddCamera(c *gin.Context) {
	var cam Camera
	if err := c.ShouldBindJSON(&cam); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	DB.Create(&cam)
	recordQueue <- &cam
	c.JSON(http.StatusCreated, cam)
}

func DeleteCamera(c *gin.Context) {
	id := c.Param("id")
	camID, _ := strconv.Atoi(id)
	DB.Delete(&Camera{}, camID)
	c.Status(http.StatusNoContent)
}
 